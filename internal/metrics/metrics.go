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

	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total number of API requests",
		},
		[]string{"updateType", "updateSource"},
	)

	latencyHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Request latency in seconds",
			Buckets: []float64{0.1, 0.3, 0.5, 0.8, 1, 2, 5},
		},
		[]string{},
	)

	responseStatusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_response_status_count",
			Help: "Count of response status codes",
		},
		[]string{"status"},
	)

	rpsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "api_requests_per_second",
			Help: "Requests per second",
		},
	)

	throughputCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "api_response_throughput_bytes_total",
			Help: "Total response throughput in bytes",
		},
	)
)

func Init() {
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		requestsTotal,
		latencyHistogram,
		responseStatusCounter,
		rpsGauge,
		throughputCounter,
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

		if rw.status != http.StatusOK {
			logger.ZapLogger.Warn("update type or source not defined")
			return
		}

		requestsTotal.WithLabelValues(updateType, updateSource).Inc()
		latencyHistogram.WithLabelValues().Observe(duration)

		statusCategory := strconv.Itoa(rw.status/100) + "xx"
		responseStatusCounter.WithLabelValues(statusCategory).Inc()

		rpsGauge.Set(float64(1))

		contentLength := w.Header().Get("Content-Length")
		if size, err := strconv.Atoi(contentLength); err == nil {
			throughputCounter.Add(float64(size))
		}

		logger.ZapLogger.Info("update processed",
			zap.String("duration", strconv.FormatFloat(duration, 'f', -1, 64)),
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
