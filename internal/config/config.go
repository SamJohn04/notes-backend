package config

import "os"

type Config struct {
	ServerPort string
}

func Load() *Config {
	return &Config{
		ServerPort: getEnv("PORT", "8000"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
