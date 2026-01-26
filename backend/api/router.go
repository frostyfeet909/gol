package api

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

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// if h.Config.Debug {
	// 	gin.SetMode(gin.DebugMode)
	// } else {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	registerSwagger(router)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	users.RegisterRouter(v1, h.UserHandler)

	return router
}
