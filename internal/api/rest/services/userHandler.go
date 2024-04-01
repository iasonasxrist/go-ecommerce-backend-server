package services

import (
	"ecommerce.com/config"
	"ecommerce.com/internal/api/rest"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type UserHandler struct {
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	var logger = slog.Default()

	handler := &UserHandler{}

	app.POST("/register", handler.Register)
	app.POST("/login", handler.Login)

	app.GET("/verify", handler.Verify)
	app.POST("/verify", handler.Verify)
	app.GET("/profile", handler.GetProfile)
	app.POST("/profile", handler.CreateProfile)

	app.GET("/cart", handler.Cart)
	app.POST("/cart", handler.Cart)
	app.GET("/order/:id", handler.Order)
	app.POST("/order", handler.Order)

	app.POST("/become-seller", handler.BecomeSeller)

	// app.Run(config.ServerPort)

}

func (h *UserHandler) Register(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "register"})
	return
}

func (h *UserHandler) Login(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "register"})
	return
}

func (h *UserHandler) Verify(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "register"})
	return
}

func (h *UserHandler) GetProfile(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "register"})
	return
}

func (h *UserHandler) CreateProfile(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "register"})
	return
}
