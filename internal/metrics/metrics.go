package metrics

import (
	"api-gateway/internal/logger"
	"api-gateway/internal/models"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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

	// Request metrics
	requestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total number of API requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	requestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "API request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Update metrics
	updatesTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "telegram_updates_total",
			Help: "Total number of Telegram updates received",
		},
		[]string{"update_type", "source"},
	)

	unknownUpdates = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "telegram_unknown_updates_total",
			Help: "Total number of unknown update types or sources",
		},
		[]string{"update_type", "source"},
	)

	noSuitableModule = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "telegram_no_suitable_module_total",
			Help: "Total number of updates with no suitable module",
		},
		[]string{"update_type", "source"},
	)

	// Module metrics
	moduleRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "module_requests_total",
			Help: "Total number of requests sent to modules",
		},
		[]string{"module", "transport_type", "status"},
	)

	moduleLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "module_request_duration_seconds",
			Help:    "Module request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"module", "transport_type"},
	)

	// Error metrics
	errorsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_errors_total",
			Help: "Total number of API errors",
		},
		[]string{"type"},
	)

	// Processing metrics
	processingStatus = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "update_processing_status",
			Help: "Status of update processing (1 for success, 0 for failure)",
		},
		[]string{"update_type", "source", "module"},
	)

	// Queue metrics
	queueSize = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "update_queue_size",
			Help: "Current size of the update processing queue",
		},
	)

	queueLatency = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "update_queue_latency_seconds",
			Help:    "Time updates spend in queue",
			Buckets: prometheus.DefBuckets,
		},
	)

	// Transport metrics
	grpcRpsGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "grpc_requests_per_second",
			Help: "gRPC requests per second",
		},
	)

	httpRpsGauge = promauto.NewGauge(
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
		requestDuration,
		updatesTotal,
		unknownUpdates,
		noSuitableModule,
		moduleRequests,
		moduleLatency,
		errorsTotal,
		processingStatus,
		queueSize,
		queueLatency,
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
		moduleName := data.moduleName
		transport := data.transport

		logger.ZapLogger.Info("Duration", zap.Float64("duration", duration))

		requestsTotal.WithLabelValues("POST", "/webhook", strconv.Itoa(rw.status)).Inc()
		requestDuration.WithLabelValues("POST", "/webhook").Observe(duration)

		statusCategory := strconv.Itoa(rw.status/100) + "xx"

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
	})
}

func SetUpdateMetrics(r *http.Request, updateType, updateSource string, forwardedModules []models.ModuleInfo, failedModules []string) {
	if data, ok := r.Context().Value(MetricsDataKey).(*metricsData); ok {
		data.updateType = updateType
		data.updateSource = updateSource

		// Логируем успешные модули с их типами транспорта
		for _, module := range forwardedModules {
			logger.ZapLogger.Info("Forwarded module",
				zap.String("module", module.ModuleName),
				zap.String("transport", module.TransportType),
			)
		}
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

// RecordRequestDuration records the duration of an API request
func RecordRequestDuration(duration float64) {
	requestDuration.WithLabelValues("POST", "/webhook").Observe(duration)
}

// RecordError records an error occurrence
func RecordError(errorType string) {
	errorsTotal.WithLabelValues(errorType).Inc()
}

// RecordUnknownUpdate records an unknown update type or source
func RecordUnknownUpdate(updateType, source string) {
	unknownUpdates.WithLabelValues(updateType, source).Inc()
}

// RecordNoSuitableModule records when no suitable module is found
func RecordNoSuitableModule(updateType, source string) {
	noSuitableModule.WithLabelValues(updateType, source).Inc()
}

// RecordModuleRequest records a request to a module
func RecordModuleRequest(module, transportType string, duration float64, err error) {
	status := "success"
	if err != nil {
		status = "error"
	}
	moduleRequests.WithLabelValues(module, transportType, status).Inc()
	moduleLatency.WithLabelValues(module, transportType).Observe(duration)
}

// RecordUpdateProcessing records the processing status of an update
func RecordUpdateProcessing(updateType, source string, forwardedModules []models.ModuleInfo, failedModules []string) {
	// Record successful modules
	for _, module := range forwardedModules {
		processingStatus.WithLabelValues(updateType, source, module.ModuleName).Set(1)
	}

	// Record failed modules
	for _, module := range failedModules {
		processingStatus.WithLabelValues(updateType, source, module).Set(0)
	}

	// Record total updates
	updatesTotal.WithLabelValues(updateType, source).Inc()
}

// RecordQueueSize records the current size of the update queue
func RecordQueueSize(size int) {
	queueSize.Set(float64(size))
}

// RecordQueueLatency records the time an update spends in the queue
func RecordQueueLatency(duration time.Duration) {
	queueLatency.Observe(duration.Seconds())
}
