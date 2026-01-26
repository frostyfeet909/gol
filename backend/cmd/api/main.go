package main

import (
	"api/api"
	"api/config"
	"context"
	"log/slog"
)

func main() {
	cfg := new(config.Config)
	cfg.LoadConfig()

	if cfg.Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	eng, cleanup := api.Build(context.Background(), *cfg)

	defer cleanup()
	eng.Run(":8080")

}
