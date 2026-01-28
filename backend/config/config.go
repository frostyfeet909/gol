package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Debug   bool   `env:"DEBUG"                 envDefault:"true"`
	AppPort int    `env:"APP_PORT"              envDefault:"8080"`
	DBURL   string `env:"DB_URL,unset,required"`
	Version string `env:"VERSION"               envDefault:"v0.0.1"`
}

func (c *Config) LoadConfig() {
	loadDotenv()
	if err := env.Parse(c); err != nil {
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
		logger.Error(fmt.Sprintf("%s: %s", "Failed to parse config from env", err))
		os.Exit(1)
	}
}
