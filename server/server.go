package server

import (
	"marine-backend/server/handles"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://example.com"}, // Allow specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                // Allow specific methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},     // Allow specific headers
		AllowCredentials: true,                                                    // Allow credentials (cookies, authorization headers)
		MaxAge:           12 * 60 * 60,                                            // Cache preflight response for 12 hours
	}))

	e.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Health!")
	})
	v1 := e.Group("/api/v1")

	port := v1.Group("/port")
	port.GET("", handles.GetPortByCode)                       // 获取单个Port
	port.POST("", handles.GetPortsByCode)                     // 批量获取Ports
	port.POST("/traffic", handles.GetPortTrafficByMonth)      // 获取逐月交通记录
	port.GET("/throughput", handles.GetPortThroughputByMonth) // 获取逐月吞吐量
}
