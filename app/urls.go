package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vakil-bashi/log-system/controllers/logs"
	"github.com/vakil-bashi/log-system/controllers/ping"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func urlPatterns() {
	router.GET("/ping", ping.Ping)

	router.POST("/logs/insert", logs.InsertLogs)
}
