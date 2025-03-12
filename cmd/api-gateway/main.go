package main

import (
	telegram_api "api-gateway/internal/generated/telegram-api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiGateway struct {
		Port     int    `yaml:"port"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"api-gateway"`
	BotModules map[string]struct {
		Grpc struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"grpc"`
		AllowedUpdates map[string][]string `yaml:"allowed-updates"`
	} `yaml:"bot-modules"`
	Logs struct {
		Address string `yaml:"address"`
	} `yaml:"where-to-send-logs"`
	Metrics struct {
		Address string `yaml:"address"`
	} `yaml:"where-to-send-metrics"`
}

var (
	config Config
	logger *zap.Logger

	requestCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "api_gateway_requests_total",
			Help: "Total number of received requests",
		},
	)

	requestsByChatType = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_gateway_requests_by_chat_type",
			Help:    "Requests per second by chat type",
			Buckets: prometheus.ExponentialBuckets(1, 2, 10),
		},
		[]string{"chat_type"},
	)

	requestsByUpdateType = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_gateway_requests_by_update_type",
			Help:    "Requests per second by update type",
			Buckets: prometheus.ExponentialBuckets(1, 2, 10),
		},
		[]string{"update_type"},
	)
)

type UpdateSource string

const (
	UpdateSourceBusinessAccount UpdateSource = "BusinessAccount"
	UpdateSourceChannel         UpdateSource = "Channel"
	UpdateSourceGroup           UpdateSource = "Group"
	UpdateSourceSuperGroup      UpdateSource = "SuperGroup"
	UpdateSourcePrivateChat     UpdateSource = "PrivateChat"
	UpdateSourceInlineMode      UpdateSource = "InlineMode"
	UpdateSourcePoll            UpdateSource = "Poll"
	UpdateSourcePayment         UpdateSource = "Payment"
	UpdateSourceUnknown         UpdateSource = "Unknown"
)

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

func init() {
	file, err := os.ReadFile("gateway-conf.yml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	logger, _ = zap.NewProduction()
	defer logger.Sync()

	prometheus.MustRegister(requestCount, requestsByChatType, requestsByUpdateType)
}

func main() {
	http.HandleFunc("/"+config.ApiGateway.Endpoint, handleUpdate)
	http.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf(":%d", config.ApiGateway.Port)
	logger.Info("Starting API Gateway", zap.String("port", strconv.Itoa(config.ApiGateway.Port)))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	var update telegram_api.Update

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		logger.Error("Failed to decode JSON", zap.Error(err))
		return
	}

	// Логируем полученный update
	logger.Info("Received update", zap.Any("update", update))
}

func extractChatType(update map[string]interface{}) (string, bool) {
	if msg, ok := update["message"].(map[string]interface{}); ok {
		if chat, ok := msg["chat"].(map[string]interface{}); ok {
			if chatType, ok := chat["type"].(string); ok {
				return chatType, true
			}
		}
	}
	return "", false
}

func extractUpdateType(update map[string]interface{}) (string, bool) {
	for key := range update {
		if key != "update_id" {
			return key, true
		}
	}
	return "", false
}

func dispatchUpdate(chatType, updateType string, update map[string]interface{}) {
	for module, settings := range config.BotModules {
		if updates, ok := settings.AllowedUpdates[chatType]; ok {
			for _, allowedUpdate := range updates {
				if allowedUpdate == updateType {
					go sendToModule(module)
					return
				}
			}
		}
	}
}

func sendToModule(module string) {
	logger.Info("Sent update to module", zap.String("module", module))
}
