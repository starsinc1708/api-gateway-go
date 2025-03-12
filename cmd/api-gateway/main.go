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
	UpdateSourceBusinessAccount UpdateSource = "business_account"
	UpdateSourceChannel         UpdateSource = "channel"
	UpdateSourceGroup           UpdateSource = "group"
	UpdateSourceSuperGroup      UpdateSource = "super_group"
	UpdateSourcePrivateChat     UpdateSource = "private_chat"
	UpdateSourceInlineMode      UpdateSource = "inline_mode"
	UpdateSourcePayment         UpdateSource = "payment"
	UpdateSourceUnknown         UpdateSource = "unknown"
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
		http.Error(w, "Invalid request", http.StatusAccepted)
		return
	}

	updateType, _ := extractUpdateType(update)
	//if !ok {
	//	http.Error(w, "Unknown update type", http.StatusAccepted)
	//	return
	//}

	updateSource, _ := extractUpdateSource(update, updateType)
	//if !ok {
	//	http.Error(w, "Unknown update source", http.StatusAccepted)
	//	return
	//}

	logger.Info("Received update", zap.Any("update_type", updateType), zap.Any("update_source", updateSource))

	requestsByUpdateType.WithLabelValues(updateType).Observe(1)

}

func extractUpdateSource(update telegram_api.Update, updateType string) (UpdateSource, bool) {
	if updateType == "business_connection" || updateType == "deleted_business_messages" || updateType == "edited_business_message" {
		return UpdateSourceBusinessAccount, true
	}
	if updateType == "callback_query" {
		chatType := update.CallbackQuery.Message.GetInaccessibleMessage().Chat.Type
		if chatType == "" {
			chatType = update.CallbackQuery.Message.GetMessage().Chat.Type
		}
		return getSourceFromChatType(chatType), true
	}
	if updateType == "channel_post" || updateType == "edited_channel_post" {
		return UpdateSourceChannel, true
	}
	if updateType == "chat_boost" {
		return getSourceFromChatType(update.ChatBoost.Chat.Type), true
	}
	if updateType == "removed_chat_boost" {
		return getSourceFromChatType(update.RemovedChatBoost.Chat.Type), true
	}
	if updateType == "chat_member" {
		return getSourceFromChatType(update.ChatMember.Chat.Type), true
	}
	if updateType == "chat_join_request" {
		return getSourceFromChatType(update.ChatJoinRequest.Chat.Type), true
	}
	if updateType == "my_chat_member" {
		return getSourceFromChatType(update.MyChatMember.Chat.Type), true
	}
	if updateType == "chosen_inline_result" || updateType == "inline_query" {
		return UpdateSourceInlineMode, true
	}
	if updateType == "message" {
		return getSourceFromChatType(update.Message.Chat.Type), false
	}
	if updateType == "edited_message" {
		return getSourceFromChatType(update.EditedMessage.Chat.Type), false
	}

	/*if updateType == "poll" || updateType == "poll_answer" {
		return getSourceFromChatType(update.Poll.)
	}*/

	return "", false
}

func getSourceFromChatType(chatType string) UpdateSource {
	switch chatType {
	case "group":
		return UpdateSourceGroup
	case "super_group":
		return UpdateSourceSuperGroup
	case "private":
		return UpdateSourcePrivateChat
	case "channel":
		return UpdateSourceChannel
	default:
		return UpdateSourceUnknown
	}
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
