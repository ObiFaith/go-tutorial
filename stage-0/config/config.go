package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	MODE string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		PORT: getEnv("PORT", "8000"),
		MODE: os.Getenv("GIN_MODE"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key);
	if value == ""{
		return fallback
	}
	return  value
}
