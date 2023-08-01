package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vakil-bashi/log-system/app"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app.StartApplication()
}
