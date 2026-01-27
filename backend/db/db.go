package db

import (
	"api/config"
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	dbpool, err := pgxpool.New(ctx, cfg.DBURL)

	if err != nil {
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
		logger.ErrorContext(ctx, "failed to connect to database", slog.Any("err", err))
		os.Exit(1)
	}

	return dbpool
}
