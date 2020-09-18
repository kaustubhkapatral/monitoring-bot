package config

import (
	"log"
	"os"
)

type AppConfig struct {
	Hex     string
	Token   string
	LCD     string
	ChatId  string
	ValAddr string
}

var NewApp AppConfig

func SetConfig() {
	NewApp = AppConfig{
		Hex:     getEnv("Hex", ""),
		ValAddr: getEnv("ValAddr", ""),
		LCD:     getEnv("LCD", ""),
		Token:   getEnv("Token", ""),
		ChatId:  getEnv("ChatId", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else if defaultVal == "" {
		log.Fatalf("Environment variable %s cannot have a nil value", key)
	}
	return defaultVal
}
