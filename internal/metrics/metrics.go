package metrics

import (
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type contextKey string

const (
	UpdateTypeKey   contextKey = "updateType"
	UpdateSourceKey contextKey = "updateSource"
)

var (
	RequestsByUpdateType = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_gateway_requests_by_update_type_total",
			Help: "Total number of requests by update type",
		},
		[]string{"update_type"},
	)

	RequestsByUpdateSource = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_gateway_requests_by_update_source_total",
			Help: "Total number of requests by update source",
		},
		[]string{"update_source"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_gateway_request_duration_seconds",
			Help:    "Duration of request processing",
			Buckets: []float64{0.001, 0.01, 0.1, 0.3, 0.5, 1, 2, 5},
		},
		[]string{"update_type", "update_source"},
	)

	ResponseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_gateway_response_status_total",
			Help: "Total number of responses by status code",
		},
		[]string{"code"},
	)
)

// Новый реестр Prometheus (чтобы избежать дубликатов)
var (
	once     sync.Once
	registry = prometheus.NewRegistry()
)

// RegisterMetrics - регистрирует метрики (только один раз)
func RegisterMetrics() {
	once.Do(func() {
		// Используем новый реестр вместо глобального
		registry.MustRegister(
			RequestsByUpdateType,
			RequestsByUpdateSource,
			RequestDuration,
			ResponseStatus,
			collectors.NewGoCollector(), // системные метрики Go
		)
		zap.L().Info("Prometheus metrics registered successfully")
	})
}

// Handler возвращает HTTP-обработчик для Prometheus
func Handler() http.Handler {
	return promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()
		updateType, _ := r.Context().Value(UpdateTypeKey).(string)
		updateSource, _ := r.Context().Value(UpdateSourceKey).(string)

		if updateType == "" {
			updateType = "unknown"
		}
		if updateSource == "" {
			updateSource = "unknown"
		}

		// Логирование обновления метрик
		zap.L().Info("Updating Prometheus metrics",
			zap.String("update_type", updateType),
			zap.String("update_source", updateSource),
			zap.Float64("duration", duration),
			zap.Int("status", rw.status),
		)

		RequestDuration.WithLabelValues(updateType, updateSource).Observe(duration)
		RequestsByUpdateType.WithLabelValues(updateType).Inc()
		RequestsByUpdateSource.WithLabelValues(updateSource).Inc()
		ResponseStatus.WithLabelValues(http.StatusText(rw.status)).Inc()
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
