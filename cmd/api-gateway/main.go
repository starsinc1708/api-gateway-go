package main

import (
	telegram_api "api-gateway/internal/generated/telegram-api"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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

	prometheus.MustRegister(requestCount, requestsByUpdateType)
}

func main() {
	http.HandleFunc("/"+config.ApiGateway.Endpoint, handleUpdate)
	http.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf(":%d", config.ApiGateway.Port)
	logger.Info("Starting API Gateway", zap.String("port", strconv.Itoa(config.ApiGateway.Port)))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	requestCount.Inc()

	var update telegram_api.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updateType, ok := extractUpdateType(update)
	if !ok {
		http.Error(w, "Unknown chat type", http.StatusBadRequest)
		return
	}

	logger.Info("Received update", zap.Any("update_type", updateType))

	requestsByUpdateType.WithLabelValues(updateType).Observe(1)

}

func extractUpdateType(update telegram_api.Update) (string, bool) {
	v := reflect.ValueOf(update)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() != reflect.Pointer {
			continue
		}
		if !field.IsNil() {
			fieldName := v.Type().Field(i).Name
			if fieldName != "" {
				return toSnakeCase(fieldName), true
			}
		}
	}
	return "", false
}

func toSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
