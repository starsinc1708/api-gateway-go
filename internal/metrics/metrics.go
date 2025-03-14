package metrics

import (
	"api-gateway/internal/logger"
	"context"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	moduleName   string
	transport    string
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

	// New metrics
	grpcRpsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "grpc_requests_per_second",
			Help: "gRPC requests per second",
		},
	)

	httpRpsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_per_second",
			Help: "HTTP requests per second",
		},
	)

	requestsPerModule = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_per_module",
			Help: "Total requests forwarded to each bot module",
		},
		[]string{"module"},
	)

	moduleRequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "module_request_latency_seconds",
			Help:    "Latency of requests to bot modules",
			Buckets: []float64{0.1, 0.3, 0.5, 0.8, 1, 2, 5},
		},
		[]string{"module"},
	)

	moduleResponseStatusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "module_response_status_count",
			Help: "Count of response status codes from bot modules",
		},
		[]string{"module", "status"},
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
		grpcRpsGauge,
		httpRpsGauge,
		requestsPerModule,
		moduleRequestLatency,
		moduleResponseStatusCounter,
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
		moduleName := data.moduleName
		transport := data.transport

		if rw.status != http.StatusOK {
			logger.ZapLogger.Warn("update type or source not defined")
			return
		}

		requestsTotal.WithLabelValues(updateType, updateSource).Inc()
		latencyHistogram.WithLabelValues().Observe(duration)

		statusCategory := strconv.Itoa(rw.status/100) + "xx"
		responseStatusCounter.WithLabelValues(statusCategory).Inc()

		rpsGauge.Set(float64(1))

		// Transport-specific metrics
		if transport == "gRPC" {
			grpcRpsGauge.Inc()
		} else if transport == "HTTP" {
			httpRpsGauge.Inc()
		}

		// Module-specific metrics
		if moduleName != "" {
			requestsPerModule.WithLabelValues(moduleName).Inc()
			moduleRequestLatency.WithLabelValues(moduleName).Observe(duration)
			moduleResponseStatusCounter.WithLabelValues(moduleName, statusCategory).Inc()
		}

		contentLength := w.Header().Get("Content-Length")
		if size, err := strconv.Atoi(contentLength); err == nil {
			throughputCounter.Add(float64(size))
		}
	})
}

func SetUpdateMetrics(r *http.Request, updateType, updateSource, moduleName, transport string) {
	if data, ok := r.Context().Value(MetricsDataKey).(*metricsData); ok {
		data.updateType = updateType
		data.updateSource = updateSource
		data.moduleName = moduleName
		data.transport = transport
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
