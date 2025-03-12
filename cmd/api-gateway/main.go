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
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("gateway-conf.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация логгера
	logger.InitLogger(cfg.Logs.Address)
	defer logger.Logger.Sync()

	// Инициализация метрик (гарантированно один раз)
	metrics.InitMetrics()

	// Регистрация обработчика с middleware для метрик
	http.Handle("/"+cfg.ApiGateway.Endpoint, metrics.MetricsMiddleware(http.HandlerFunc(handlers.HandleUpdate)))

	// Регистрация эндпоинта для Prometheus метрик
	http.Handle("/metrics", metrics.Handler())

	// Запуск сервера
	addr := fmt.Sprintf(":%d", cfg.ApiGateway.Port)
	logger.Logger.Info("Starting API Gateway", zap.String("port", addr))
	log.Fatal(http.ListenAndServe(addr, nil))
}
