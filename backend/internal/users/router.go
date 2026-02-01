package users

import "github.com/labstack/echo/v5"

func RegisterRouter(rg *echo.Group, h *Handler) {
	users := rg.Group("/users")
	users.GET("", h.list)
	users.GET("/:id", h.get)
	users.POST("", h.create)
	users.DELETE("/:id", h.delete)
}
