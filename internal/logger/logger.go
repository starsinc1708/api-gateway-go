package logger

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"strings"
)

var Logger *zap.Logger

func InitLogger(logsAddress string) {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	var cores []zapcore.Core
	cores = append(cores, consoleCore) // Всегда логируем в консоль

	if logsAddress != "" {
		if !(strings.HasPrefix(logsAddress, "http://") || strings.HasPrefix(logsAddress, "https://")) {
			logsAddress = "http://" + logsAddress // Добавляем схему, если её нет
		}
		httpEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		httpCore := zapcore.NewCore(httpEncoder, zapcore.AddSync(&httpPostSyncer{url: logsAddress}), zapcore.InfoLevel)
		cores = append(cores, httpCore)
	}

	core := zapcore.NewTee(cores...) // Теперь логируем в нужные места
	Logger = zap.New(core, zap.AddCaller())
}

type httpPostSyncer struct {
	url string
}

func (h *httpPostSyncer) Write(p []byte) (n int, err error) {
	resp, err := http.Post(h.url, "application/json", bytes.NewBuffer(p))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return len(p), nil
}

func (h *httpPostSyncer) Sync() error {
	return nil
}
