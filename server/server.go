package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	e.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Health!")
	})
}
