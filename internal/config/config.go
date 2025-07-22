package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerPort string
	JWTSecret  string
}

var Cfg Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Ignore this message if not needed.")
	}

	Cfg = Config{
		ServerPort: getEnv("PORT", "8000"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
