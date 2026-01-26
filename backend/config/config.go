package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Debug   bool   `env:"DEBUG" envDefault:"true"`
	AppPort int    `env:"APP_PORT" envDefault:"8080"`
	DbURL   string `env:"DB_URL,unset,required"`
}

func (c *Config) LoadConfig() {
	loadDotenv()
	if err := env.Parse(c); err != nil {
		slog.Error(fmt.Sprintf("%s: %s", "Failed to parse config from env", err))
		os.Exit(1)
	}
}
