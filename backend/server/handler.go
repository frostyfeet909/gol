package server

import (
	"api/config"
	"api/db"
	"api/internal/users"
	"context"

	"github.com/labstack/echo/v5"
)

type Cleanup func()

func Build(ctx context.Context, cfg config.Config) (*echo.Echo, Cleanup) {
	db := db.Connect(ctx, &cfg)

	userRepo := users.NewPostgresRepo(db)
	userSvc := users.NewService(userRepo)
	userHandler := users.NewHandler(userSvc)

	app := createRouters(Handlers{
		Config:      cfg,
		UserHandler: userHandler,
	})

	cleanup := func() {
		db.Close()
	}

	return app, cleanup
}
