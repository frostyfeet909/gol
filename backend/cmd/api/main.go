package main

import (
	"api/config"
	"api/server"
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

	eng, cleanup := server.Build(context.Background(), *cfg)

	defer cleanup()
	if err := eng.Start(":8080"); err != nil {
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
		logger.Error("Failed to run server", "error", err)
	}
}
