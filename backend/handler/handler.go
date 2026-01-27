package handler

import (
	"api/config"
	"api/db"
	"api/internal/users"
	"context"

	"github.com/gin-gonic/gin"
)

type Cleanup func()

func Build(ctx context.Context, cfg config.Config) (*gin.Engine, Cleanup) {
	db := db.Connect(ctx, &cfg)

	userRepo := users.NewPostgresRepo(db)
	userSvc := users.NewService(userRepo)
	userHandler := users.NewHandler(userSvc)

	app := CreateRouters(Handlers{
		Config:      cfg,
		UserHandler: userHandler,
	})

	cleanup := func() {
		db.Close()
	}

	return app, cleanup
}
