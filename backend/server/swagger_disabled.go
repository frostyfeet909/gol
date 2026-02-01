//go:build !swagger

package server

import "github.com/labstack/echo/v5"

func registerSwagger(_ *echo.Echo) {}
