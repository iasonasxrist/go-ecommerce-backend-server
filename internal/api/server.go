package api

import (
	"log"

	"ecommerce.com/config"
	"ecommerce.com/helper"
	"ecommerce.com/internal/api/rest"
	"ecommerce.com/internal/api/rest/services"
	"ecommerce.com/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {

	app := gin.Default()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("connection error %v\n", err)
	}
	log.Printf("database connected")
	log.Print(db)

	//run migration
	db.AutoMigrate(&domain.User{})

	auth := helper.SetupAuth(config.AppSecret)

	restHandler := &rest.RestHandler{
		App:  app,
		Db:   db,
		Auth: auth,
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
