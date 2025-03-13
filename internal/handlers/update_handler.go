package handlers

import (
	"api-gateway/internal/generated/telegram-api"
	"api-gateway/internal/logger"
	"api-gateway/internal/metrics"
	"api-gateway/internal/services"
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
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
		logger.Logger.Error("Failed to extract update source",
			zap.String("update_type", updateType),
			zap.Any("update", update))
		http.Error(w, "Failed to extract update source", http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), metrics.UpdateTypeKey, updateType)
	ctx = context.WithValue(ctx, metrics.UpdateSourceKey, updateSource)
	*r = *r.WithContext(ctx)

	logger.Logger.Info("Processing update",
		zap.String("update_type", updateType),
		zap.String("update_source", string(updateSource)),
	)

	w.WriteHeader(http.StatusOK)
}
