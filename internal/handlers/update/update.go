package update

import (
	"api-gateway/internal/config"
	"api-gateway/internal/logger"
	"api-gateway/internal/metrics"
	"api-gateway/internal/models"
	"api-gateway/internal/models/telegram"
	"api-gateway/internal/services"
	"api-gateway/internal/transport"
	"api-gateway/internal/utils"
	"encoding/json"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime).Seconds()
		metrics.RecordRequestDuration(duration)
	}()

	var update telegram.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		logger.ZapLogger.Error("Failed to decode request body", zap.Error(err))
		metrics.RecordError("decode_request")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	logger.ZapLogger.Info("Received Telegram update",
		zap.Int64("update_id", update.UpdateID),
		zap.String("update_type", services.ExtractUpdateType(update)),
	)

	updateType := services.ExtractUpdateType(update)
	source, fromId := services.ExtractUpdateSource(update, updateType)
	updateSource := string(source)

	if updateType == "" || updateSource == "" {
		logger.ZapLogger.Warn("Unknown update type or source",
			zap.String("updateType", updateType),
			zap.String("updateSource", updateSource),
		)
		metrics.RecordUnknownUpdate(updateType, updateSource)
		w.WriteHeader(http.StatusOK)
		return
	}

	suitableModules := utils.ExtractSuitableModules(h.cfg.BotModules.Modules, updateSource, updateType)
	if len(suitableModules) == 0 {
		logger.ZapLogger.Warn("No suitable bot module found for update",
			zap.String("updateType", updateType),
			zap.String("updateSource", updateSource),
			zap.Int64("From", fromId),
		)
		metrics.RecordNoSuitableModule(updateType, updateSource)
		w.WriteHeader(http.StatusOK)
		return
	}

	var (
		forwardedModules []models.ModuleInfo
		failedModules    []string
	)

	for name, module := range suitableModules {
		var errTransport error
		var transportType string
		moduleStartTime := time.Now()

		if module.GRPC.Host != "" {
			transportType = "gRPC"
			errTransport = transport.SendGrpc(module.GRPC.Host, module.GRPC.Port, update, updateType, updateSource, fromId)
		} else if module.HTTP.Host != "" {
			transportType = "HTTP"
			errTransport = transport.SendHttp(module.HTTP.Host, module.HTTP.Port, update, updateType, updateSource, fromId)
		}

		moduleDuration := time.Since(moduleStartTime).Seconds()
		metrics.RecordModuleRequest(name, transportType, moduleDuration, errTransport)

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

	metrics.RecordUpdateProcessing(updateType, updateSource, forwardedModules, failedModules)
	w.WriteHeader(http.StatusOK)
}

func getModuleNames(modules []models.ModuleInfo) []string {
	names := make([]string, len(modules))
	for i, module := range modules {
		names[i] = module.ModuleName
	}
	return names
}
