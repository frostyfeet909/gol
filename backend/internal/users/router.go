package users

import "github.com/gin-gonic/gin"

func RegisterRouter(rg *gin.RouterGroup, h *Handler) {
	users := rg.Group("/users")
	{
		users.POST("", h.Create)
		users.GET("/:id", h.Get)
	}
}
