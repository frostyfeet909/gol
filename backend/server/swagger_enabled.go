//go:build swagger

package server

import (
	_ "api/docs"
	"net/http"

	"github.com/labstack/echo/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func registerSwagger(r *echo.Echo) {
	r.GET("/swagger/*", echo.WrapHandler(http.HandlerFunc(httpSwagger.WrapHandler)))
}
