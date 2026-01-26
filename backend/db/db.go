package db

import (
	"api/config"
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	dbpool, err := pgxpool.New(ctx, cfg.DbURL)

	if err != nil {
		slog.Error(fmt.Sprintf("%s: %s", "Failed to connect to database", err))
		os.Exit(1)
	}

	return dbpool
}
