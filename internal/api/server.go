package api

import (
	"ecommerce.com/config"
	"github.com/gin-gonic/gin"
)

func StartServer(config config.AppConfig) {

	app := gin.Default()

	app.GET("/health", HealthCheck)
	app.Run(config.ServerPort)
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "Ok"})
}
