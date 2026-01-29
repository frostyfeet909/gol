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
	router := gin.Default()

	router.Use(ginhelmet.Default())

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
