//go:build dotenv

package config

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func loadDotenv() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found")
	}
}
