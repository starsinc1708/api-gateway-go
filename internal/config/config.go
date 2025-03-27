package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiGateway ApiGatewayConfig `yaml:"api-gateway"`
	Logs       LogConfig        `yaml:"logs"`
	Metrics    MetricsConfig    `yaml:"metrics"`
	BotModules BotModulesConfig `yaml:"bot-modules"`
}

type ApiGatewayConfig struct {
	Port     int    `yaml:"port"`
	Endpoint string `yaml:"endpoint"`
}

type LogConfig struct {
	Address string `yaml:"address"`
	Level   string `yaml:"level"`
}

type MetricsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
	Port    int    `yaml:"port"`
}

type BotModulesConfig struct {
	Modules map[string]ModuleConfig `yaml:"modules"`
}

type ModuleConfig struct {
	GRPC GRPCConfig `yaml:"grpc"`
	HTTP HTTPConfig `yaml:"http"`
}

type GRPCConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type HTTPConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func LoadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}
