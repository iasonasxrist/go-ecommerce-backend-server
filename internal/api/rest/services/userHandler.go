package services

import (
	"ecommerce.com/internal/api/rest"
	"ecommerce.com/internal/dto"
	"ecommerce.com/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	// add and inject userService into handler
	svc := service.UserService{}
	handler := UserHandler{
		svc: svc,
	}

	app.POST("/register", handler.Register)
	app.POST("/login", handler.Login)

	app.GET("/verify", handler.GetVerificationCode)
	app.POST("/verify", handler.Verify)
	app.GET("/profile", handler.GetProfile)
	app.POST("/profile", handler.CreateProfile)

	app.GET("/cart", handler.GetCart)
	app.POST("/cart", handler.AddToCart)
	app.GET("/order/:id", handler.GetOrder)
	app.POST("/order", handler.GetOrders)

	app.POST("/become-seller", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *gin.Context) {

	var user dto.UserSignup

	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "plese provide valid outputs"})
		return
	}
	token, err := h.svc.Signup(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}
	ctx.JSON(200, gin.H{"message": "register"})
	return

	ctx.JSON(http.StatusOK, gin.H{"message": token})
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

func (h *UserHandler) GetVerificationCode(ctx *gin.Context) {

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
	ctx.JSON(200, gin.H{"message": "Get Profile"})
	return
}

func (h *UserHandler) CreateProfile(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "Create Profile"})
	return
}

func (h *UserHandler) GetCart(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "Get Cart"})
	return
}

func (h *UserHandler) AddToCart(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "Add To Cart"})
	return
}

func (h *UserHandler) GetOrders(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "Get Orders"})
	return
}

func (h *UserHandler) GetOrder(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "Get Order"})
	return
}

func (h *UserHandler) BecomeSeller(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(200, gin.H{"message": "Become a Seller"})
	return
}
