package handlers

import (
	"api-gateway/internal/config"
	"api-gateway/internal/generated/telegram-api"
	"api-gateway/internal/logger"
	"api-gateway/internal/metrics"
	"api-gateway/internal/services"
	"api-gateway/internal/transport"
	"api-gateway/internal/utils"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func HandleUpdate(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var update telegram_api.Update
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			logger.ZapLogger.Error("Failed to decode request body", zap.Error(err))
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		updateType := services.ExtractUpdateType(update)
		updateSource := string(services.ExtractUpdateSource(update, updateType))

		if updateType == "" || updateSource == "" {
			logger.ZapLogger.Warn("Unknown update type or source",
				zap.String("updateType", updateType),
				zap.String("updateSource", updateSource),
			)
			http.Error(w, "Unsupported update", http.StatusBadRequest)
			return
		}

		var (
			forwarded     bool
			transportType string
			moduleName    string
			errTransport  error
		)

		for name, module := range cfg.BotModules {
			allowedTypes, exists := module.AllowedUpdates[updateSource]
			if exists && utils.Contains(allowedTypes, updateType) {
				moduleName = name
				if module.Grpc.Host != "" {
					transportType = "gRPC"
					errTransport = transport.SendGrpc(module.Grpc.Host, module.Grpc.Port, update)
				} else if module.Http.Host != "" {
					transportType = "HTTP"
					errTransport = transport.SendHttp(module.Http.Host, module.Http.Port, update)
				}
				forwarded = true
				break
			}
		}

		if !forwarded {
			logger.ZapLogger.Warn("No suitable bot module found for update",
				zap.String("updateType", updateType),
				zap.String("updateSource", updateSource),
			)
		}

		if errTransport != nil {
			logger.ZapLogger.Error("Failed to forward update",
				zap.String("module", moduleName),
				zap.String("transport", transportType),
				zap.Error(errTransport),
			)
		}

		metrics.SetUpdateMetrics(r, updateType, updateSource, moduleName, transportType)

		logger.ZapLogger.Info("Update processed",
			zap.String("updateType", updateType),
			zap.String("updateSource", updateSource),
			zap.String("module", moduleName),
			zap.String("transport", transportType),
			zap.Bool("forwarded", forwarded),
		)

		w.WriteHeader(http.StatusOK)
	}
}
