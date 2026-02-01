package server

import (
	"api/config"
	"net/http"

	"github.com/labstack/echo/v5"
)

type healthResponse struct {
	OK bool `json:"ok"`
}

type versionResponse struct {
	Version string `json:"version"`
}

// Health godoc
// @Summary      Health check
// @Description  Check if the service is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  healthResponse
// @Router       /health [get]
func health(c *echo.Context) error {
	return c.JSON(http.StatusOK, &healthResponse{OK: true})
}

// Version godoc
// @Summary      Get service version
// @Description  Fetch the current version of the service
// @Tags         health
// @Produce      json
// @Success      200  {object}  versionResponse
// @Router       /version [get]
func version(cfg config.Config) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.JSON(http.StatusOK, &versionResponse{Version: cfg.Version})
	}
}
