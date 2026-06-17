package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	GinMode        string
	GenderizeApi   string
	AllowedOrigins []string
}

var (
	globalConfig *Config
	once         sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}

		origins := os.Getenv("ALLOWED_ORIGINS")
		var allowedOrigins []string
		if origins != "" {
			allowedOrigins = strings.Split(origins, ",")
		}

		globalConfig = &Config{
			Port:           os.Getenv("PORT"),
			GinMode:        os.Getenv("GIN_MODE"),
			GenderizeApi:   os.Getenv("GENDERIZE_API"),
			AllowedOrigins: allowedOrigins,
		}
	})

	return globalConfig
}
