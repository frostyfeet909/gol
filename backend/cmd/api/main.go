package main

import (
	"api/config"
	"api/handler"
	"context"
	"log/slog"
	"os"
)

func main() {
	cfg := new(config.Config)
	cfg.LoadConfig()

	if cfg.Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	eng, cleanup := handler.Build(context.Background(), *cfg)

	defer cleanup()
	err := eng.Run(":8080")
	if err != nil {
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
		logger.Error("Failed to run server", "error", err)
	}
}
