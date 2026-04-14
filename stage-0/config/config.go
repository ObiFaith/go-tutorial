package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	GinMode string
	GenderizeApi string
  AllowedOrigins []string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	origins := os.Getenv("ALLOWED_ORIGINS")

	return &Config{
		Port: os.Getenv("PORT"),
		GinMode: os.Getenv("GIN_MODE"),
		GenderizeApi: os.Getenv("GENDERIZE_API"),
		AllowedOrigins: strings.Split(origins, ","),
	}
}
