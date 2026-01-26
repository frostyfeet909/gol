//go:build !swagger

package api

import (
	"github.com/gin-gonic/gin"
)

func registerSwagger(r *gin.Engine) {}
