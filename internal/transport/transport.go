package transport

import (
	"api-gateway/internal/generated/telegram-api"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"

	bm "api-gateway/internal/generated/bot-module"
)

func SendHttp(host string, port int, update telegram_api.Update, updateType, updateSource string) error {
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

func SendGrpc(host string, port int, update telegram_api.Update, updateType, updateSource string) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server: %w", err)
	}
	defer conn.Close()

	client := bm.NewBotModuleServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJson, err := json.Marshal(update)
	if err != nil {
		return fmt.Errorf("failed to send gRPC request: %w", err)
	}

	_, err = client.HandleUpdate(ctx, &bm.UpdateRequest{
		UpdateJson:   string(updateJson),
		UpdateType:   updateType,
		UpdateSource: updateSource,
	})
	if err != nil {
		return fmt.Errorf("failed to send gRPC request: %w", err)
	}

	return nil
}
