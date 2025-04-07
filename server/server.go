package server

import (
	"marine-backend/server/handles"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	e.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Health!")
	})
	v1 := e.Group("/api/v1")
	port := v1.Group("/port")
	port.POST("/traffic", handles.GetPortTraffic)
}
