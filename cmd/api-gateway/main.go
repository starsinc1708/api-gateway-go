package main

import (
	"fmt"
	"log"
	"net/http"

	"api-gateway/internal/config"
	"api-gateway/internal/handlers"
	"api-gateway/internal/logger"
	"api-gateway/internal/metrics"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger.InitZapLogger(cfg.Logs.Address)
	defer logger.ZapLogger.Sync()

	metrics.Init()

	http.Handle("/metrics", metrics.Handler())
	http.Handle("/"+cfg.ApiGateway.Endpoint, metrics.MetricsMiddleware(handlers.HandleUpdate(cfg)))

	addr := fmt.Sprintf(":%d", cfg.ApiGateway.Port)
	logger.ZapLogger.Info("Starting API Gateway", zap.String("port", addr))
	log.Fatal(http.ListenAndServe(addr, nil))
}
