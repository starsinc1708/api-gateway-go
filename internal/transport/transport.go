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

	pb "api-gateway/internal/generated/telegram-api"
)

func SendHttp(host string, port int, update telegram_api.Update) error {
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

func SendGrpc(host string, port int, update telegram_api.Update) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server: %w", err)
	}
	defer conn.Close()

	client := pb.NewBotModuleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.HandleUpdate(ctx, &pb.UpdateRequest{
		UpdateJson: update.String(),
	})
	if err != nil {
		return fmt.Errorf("failed to send gRPC request: %w", err)
	}

	return nil
}
