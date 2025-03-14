package config

import (
	"gopkg.in/yaml.v3"
	"os"
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
		Http struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"http"`
		AllowedUpdates map[string][]string `yaml:"allowed-updates"`
	} `yaml:"bot-modules"`
	Logs struct {
		Address string `yaml:"address"`
	} `yaml:"where-to-send-logs"`
	Metrics struct {
		Address string `yaml:"address"`
	} `yaml:"where-to-send-metrics"`
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
