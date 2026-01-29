package handler

import (
	"api/config"
	"api/internal/users"

	"github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Config      config.Config
	UserHandler *users.Handler
}

func CreateRouters(h Handlers) *gin.Engine {
	if h.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	if !h.Config.Debug {
		router.Use(ginhelmet.Default())
		router.TrustedPlatform = gin.PlatformFlyIO
	}

	api := router.Group("/api")

	api.GET("/health", health)
	api.GET("/version", version(h.Config))

	// if h.Config.Debug {
	// 	gin.SetMode(gin.DebugMode)
	// } else {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	registerSwagger(router)

	v1 := api.Group("/v1")

	users.RegisterRouter(v1, h.UserHandler)

	return router
}
