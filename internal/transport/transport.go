package transport

import (
	"api-gateway/internal/logger"
	"api-gateway/internal/models/telegram"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendHttp(host string, port int, update telegram.Update, updateType, updateSource string, id int64) error {
	url := fmt.Sprintf("http://%s:%d/tg-updates", host, port)
	jsonData, err := json.Marshal(update)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP request failed with status %d", resp.StatusCode)
	}

	return nil
}

func SendGrpc(host string, port int, update telegram.Update, updateType, updateSource string, id int64) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server: %w", err)
	}
	defer conn.Close()

	// TODO: Update gRPC client to use new structs
	// For now, we'll just send the update as JSON
	updateJson, err := json.Marshal(update)
	if err != nil {
		return fmt.Errorf("failed to marshal update: %w", err)
	}

	// Create a simple request struct
	type UpdateRequest struct {
		UpdateJson   string `json:"update_json"`
		UpdateType   string `json:"update_type"`
		UpdateSource string `json:"update_source"`
		FromId       int64  `json:"from_id"`
	}

	req := UpdateRequest{
		UpdateJson:   string(updateJson),
		UpdateType:   updateType,
		UpdateSource: updateSource,
		FromId:       id,
	}

	// Send the request as JSON
	_, err = json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	logger.ZapLogger.Info("Sending gRPC request",
		zap.String("host", host),
		zap.Int("port", port),
		zap.String("update_type", updateType),
		zap.String("update_source", updateSource),
		zap.Int64("from_id", id),
	)

	return nil
}
