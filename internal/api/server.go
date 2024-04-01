package api

import (
	"ecommerce.com/config"
	"ecommerce.com/internal/api/rest"
	"ecommerce.com/internal/api/rest/services"
	"github.com/gin-gonic/gin"
)

func StartServer(config config.AppConfig) {

	app := gin.Default()

	restHandler := &rest.RestHandler{
		App: app,
	}

	setupRoutes(restHandler)
	app.Run(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	// user
	services.SetupUserRoutes(rh)

	// transactions

	// catalog
}
