package handlers

import (
	"api-gateway/internal/generated/telegram-api"
	"api-gateway/internal/logger"
	"api-gateway/internal/metrics"
	"api-gateway/internal/services"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var update telegram_api.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		logger.ZapLogger.Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updateType := services.ExtractUpdateType(update)
	updateSource := services.ExtractUpdateSource(update, updateType)

	metrics.SetUpdateTypeAndSource(r, updateType, string(updateSource))
	if updateType == "unknown" && updateSource == "unknown" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
