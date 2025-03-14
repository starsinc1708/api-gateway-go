package utils

import (
	"api-gateway/internal/config"
	"regexp"
	"strings"
)

func ToSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func Contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func ExtractSuitableModules(modules map[string]config.BotModuleConfig, updateSource, updateType string) map[string]config.BotModuleConfig {
	suitableModules := make(map[string]config.BotModuleConfig)

	for name, module := range modules {
		if allowedTypes, exists := module.AllowedUpdates[updateSource]; exists {
			if Contains(allowedTypes, updateType) {
				suitableModules[name] = module
			}
		}
	}
	return suitableModules
}
