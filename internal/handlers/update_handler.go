package handlers

import (
	"api-gateway/internal/config"
	"api-gateway/internal/generated/telegram-api"
	"api-gateway/internal/logger"
	"api-gateway/internal/metrics"
	"api-gateway/internal/models"
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
		source, fromId := services.ExtractUpdateSource(update, updateType)

		updateSource := string(source)

		if updateType == "" || updateSource == "" {
			logger.ZapLogger.Warn("Unknown update type or source",
				zap.String("updateType", updateType),
				zap.String("updateSource", updateSource),
			)
			http.Error(w, "Unsupported update", http.StatusBadRequest)
			return
		}

		suitableModules := utils.ExtractSuitableModules(cfg.BotModules, updateSource, updateType)
		if len(suitableModules) == 0 {
			logger.ZapLogger.Warn("No suitable bot module found for update",
				zap.String("updateType", updateType),
				zap.String("updateSource", updateSource),
				zap.Int64("From", fromId),
			)
			http.Error(w, "No suitable module found", http.StatusNotFound)
			return
		}

		var (
			forwardedModules []models.ModuleInfo
			failedModules    []string
		)

		for name, module := range suitableModules {
			var errTransport error
			var transportType string

			if module.Grpc.Host != "" {
				transportType = "gRPC"
				errTransport = transport.SendGrpc(module.Grpc.Host, module.Grpc.Port, update, updateType, updateSource, fromId)
			} else if module.Http.Host != "" {
				transportType = "HTTP"
				errTransport = transport.SendHttp(module.Http.Host, module.Http.Port, update, updateType, updateSource, fromId)
			}

			if errTransport == nil {
				forwardedModules = append(forwardedModules, models.ModuleInfo{
					ModuleName:    name,
					TransportType: transportType,
				})
			} else {
				failedModules = append(failedModules, name)
				logger.ZapLogger.Error("Failed to forward update",
					zap.String("module", name),
					zap.String("transport", transportType),
					zap.Error(errTransport),
				)
			}
		}

		logger.ZapLogger.Info("Update processed",
			zap.String("updateType", updateType),
			zap.String("updateSource", updateSource),
			zap.Strings("forwardedModules", getModuleNames(forwardedModules)),
			zap.Strings("failedModules", failedModules),
		)

		metrics.SetUpdateMetrics(r, updateType, updateSource, forwardedModules, failedModules)

		// Если хотя бы один модуль обработал запрос — возвращаем 200, иначе 500
		if len(forwardedModules) > 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func getModuleNames(modules []models.ModuleInfo) []string {
	names := make([]string, len(modules))
	for i, module := range modules {
		names[i] = module.ModuleName
	}
	return names
}
