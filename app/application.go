package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(CORSMiddleware())
	urlPatterns()
	logger.Info("about to start the vakilbashi-API v-0.0.1")

	router.Run(":9091")
}
