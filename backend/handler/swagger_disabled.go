//go:build !swagger

package handler

import (
	"github.com/gin-gonic/gin"
)

func registerSwagger(_ *gin.Engine) {}
