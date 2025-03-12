package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
	"time"
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
			Buckets: []float64{0.1, 0.5, 1, 2, 5}, // Персентили 50, 90, 95, 99
		},
		[]string{"update_type", "update_source"},
	)

	RequestsPerSecond = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "api_gateway_requests_per_second",
			Help: "Number of requests per second",
		},
		[]string{"update_type", "update_source"},
	)

	goMetrics = collectors.NewGoCollector()
)

var (
	initMetricsOnce    sync.Mutex
	metricsInitialized bool
)

var registry = prometheus.DefaultRegisterer

func InitMetrics() {
	initMetricsOnce.Lock()
	defer initMetricsOnce.Unlock()

	if metricsInitialized {
		return
	}

	if _, err := registry.(*prometheus.Registry).Gather(); err == nil {
		return
	}

	prometheus.MustRegister(
		RequestsByUpdateType,
		RequestsByUpdateSource,
		RequestDuration,
		RequestsPerSecond,
		goMetrics,
	)

	metricsInitialized = true
}

func Handler() http.Handler {
	return promhttp.Handler()
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start).Seconds()
		RequestDuration.WithLabelValues("unknown", "unknown").Observe(duration)
	})
}
