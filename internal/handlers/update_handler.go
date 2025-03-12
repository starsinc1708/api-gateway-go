package handlers

import (
	"api-gateway/internal/logger"
	"encoding/json"
	"net/http"
	"time"

	"api-gateway/internal/generated/telegram-api"
	"api-gateway/internal/metrics"
	"api-gateway/internal/services"
	"go.uber.org/zap"
)

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var update telegram_api.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		logger.Logger.Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updateType, ok := services.ExtractUpdateType(update)
	if !ok {
		logger.Logger.Error("Failed to extract update type", zap.Any("update", update))
		http.Error(w, "Failed to extract update type", http.StatusBadRequest)
		return
	}

	updateSource, ok := services.ExtractUpdateSource(update, updateType)
	if !ok {
		logger.Logger.Error("Failed to extract update source", zap.String("update_type", updateType), zap.Any("update", update))
		http.Error(w, "Failed to extract update source", http.StatusBadRequest)
		return
	}

	metrics.RequestsByUpdateType.WithLabelValues(updateType).Inc()
	metrics.RequestsByUpdateSource.WithLabelValues(string(updateSource)).Inc()
	metrics.RequestsPerSecond.WithLabelValues(updateType, string(updateSource)).Inc()

	duration := time.Since(start).Seconds()
	metrics.RequestDuration.WithLabelValues(updateType, string(updateSource)).Observe(duration)

	logger.Logger.Info("processed update",
		zap.String("update_type", updateType),
		zap.String("update_source", string(updateSource)),
		zap.Float64("processing_time", duration),
	)

	w.WriteHeader(http.StatusOK)
}
