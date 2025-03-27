package logger

import (
	"bytes"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLogger *zap.Logger

func InitZapLogger(logsAddress string) {
	// Get log level from environment variable
	logLevel := zapcore.InfoLevel
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		switch strings.ToLower(level) {
		case "debug":
			logLevel = zapcore.DebugLevel
		case "warn":
			logLevel = zapcore.WarnLevel
		case "error":
			logLevel = zapcore.ErrorLevel
		}
	}

	// Configure console logging
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel)

	var cores []zapcore.Core
	cores = append(cores, consoleCore)

	// Configure HTTP logging if address is provided
	if logsAddress != "" {
		if !(strings.HasPrefix(logsAddress, "http://") || strings.HasPrefix(logsAddress, "https://")) {
			logsAddress = "http://" + logsAddress
		}
		httpEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		httpCore := zapcore.NewCore(httpEncoder, zapcore.AddSync(&httpPostSyncer{url: logsAddress}), logLevel)
		cores = append(cores, httpCore)
	}

	core := zapcore.NewTee(cores...)
	ZapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

type httpPostSyncer struct {
	url string
}

func (h *httpPostSyncer) Write(p []byte) (n int, err error) {
	resp, err := http.Post(h.url, "application/json", bytes.NewBuffer(p))
	if err != nil {
		// Log error to stderr since we can't use the logger here
		os.Stderr.WriteString("Failed to send log to HTTP endpoint: " + err.Error() + "\n")
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		os.Stderr.WriteString("HTTP logging endpoint returned non-200 status: " + resp.Status + "\n")
	}

	return len(p), nil
}

func (h *httpPostSyncer) Sync() error {
	return nil
}
