package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ApiGatewayConfig struct {
	Port     int    `yaml:"port"`
	Endpoint string `yaml:"endpoint"`
}

type GrpcConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type HttpConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type BotModuleConfig struct {
	Grpc           GrpcConfig          `yaml:"grpc"`
	Http           HttpConfig          `yaml:"http"`
	AllowedUpdates map[string][]string `yaml:"allowed-updates"`
}

type LoggingConfig struct {
	Address string `yaml:"address"`
}

type MetricsConfig struct {
	Address string `yaml:"address"`
}

type Config struct {
	ApiGateway ApiGatewayConfig           `yaml:"api-gateway"`
	BotModules map[string]BotModuleConfig `yaml:"bot-modules"`
	Logs       LoggingConfig              `yaml:"where-to-send-logs"`
	Metrics    MetricsConfig              `yaml:"where-to-send-metrics"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
