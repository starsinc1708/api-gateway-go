package metrics

import (
	"api-gateway/internal/logger"
	"context"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type contextKey string

const (
	MetricsDataKey contextKey = "metricsData"
)

type metricsData struct {
	updateType   string
	updateSource string
}

var (
	registry = prometheus.NewRegistry()
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

	RequestDurationByTypeAndSource = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_gateway_request_duration_seconds_by_type_and_source",
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

	RequestsPerSecond = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_gateway_requests_per_second_total",
			Help: "Total number of requests per second",
		},
		[]string{},
	)

	RequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_gateway_request_latency_seconds",
			Help:    "Latency of request processing",
			Buckets: []float64{0.001, 0.01, 0.1, 0.3, 0.5, 1, 2, 5}, // Определите корзины для гистограммы
		},
		[]string{"update_type", "update_source"},
	)
)

func Init() {
	registry.MustRegister(
		RequestsByUpdateType,
		RequestsByUpdateSource,
		RequestDurationByTypeAndSource,
		ResponseStatus,
		RequestsPerSecond,
		RequestLatency,
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
}

func Handler() http.Handler {
	return promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		data := &metricsData{}
		ctx := context.WithValue(r.Context(), MetricsDataKey, data)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()

		updateType := data.updateType
		updateSource := data.updateSource

		if updateType == "" {
			updateType = "unknown"
		}
		if updateSource == "" {
			updateSource = "unknown"
		}

		RequestsByUpdateType.WithLabelValues(updateType).Inc()
		RequestsByUpdateSource.WithLabelValues(updateSource).Inc()
		RequestDurationByTypeAndSource.WithLabelValues(updateType, updateSource).Observe(duration)
		RequestLatency.WithLabelValues(updateType, updateSource).Observe(duration)
		RequestsPerSecond.WithLabelValues().Inc()

		ResponseStatus.WithLabelValues(strconv.Itoa(rw.status)).Inc()

		if rw.status != 200 {
			logger.ZapLogger.Warn("update type or source not defined",
				zap.String("updateType", updateType),
				zap.String("updateSource", updateSource),
			)
			return
		}
		logger.ZapLogger.Info("update processed",
			zap.String("updateType", updateType),
			zap.String("updateSource", updateSource),
		)
	})
}

func SetUpdateTypeAndSource(r *http.Request, updateType, updateSource string) {
	if data, ok := r.Context().Value(MetricsDataKey).(*metricsData); ok {
		data.updateType = updateType
		data.updateSource = updateSource
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
