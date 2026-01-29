package handler

import (
	"api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary      Health check
// @Description  Check if the service is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]bool
// @Router       /health [get]
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// Version godoc
// @Summary      Get service version
// @Description  Fetch the current version of the service
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /version [get]
func version(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": cfg.Version})
	}
}
