package services

import (
	"net/http"

	"ecommerce.com/helper"
	"ecommerce.com/internal/api/rest"
	"ecommerce.com/internal/dto"
	"ecommerce.com/internal/repository"
	"ecommerce.com/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
	auth helper.Auth
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	// add and inject userService into handler
	svc := service.UserService{
		Repo: repository.NewRepository(rh.Db),
		Auth: rh.Auth,
	}
	handler := UserHandler{
		svc: svc,
	}
	pubRoutes := app.Group("/users")

	pubRoutes.POST("/register", handler.Register)
	pubRoutes.POST("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorization)

	pvtRoutes.GET("/verify", handler.GetVerificationCode)
	pvtRoutes.POST("/verify", handler.Verify)
	pvtRoutes.GET("/profile", handler.GetProfile)
	pvtRoutes.POST("/profile", handler.CreateProfile)

	pvtRoutes.GET("/cart", handler.GetCart)
	pvtRoutes.POST("/cart", handler.AddToCart)
	pvtRoutes.GET("/order/:id", handler.GetOrder)
	pvtRoutes.POST("/order", handler.GetOrders)

	pvtRoutes.POST("/become-seller", handler.BecomeSeller)

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

	ctx.JSON(http.StatusOK, gin.H{"message": token})
	return

}

func (h *UserHandler) Login(ctx *gin.Context) {

	var loginInput dto.UserLogin

	err := ctx.BindJSON(&loginInput)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "plese provide valid inputs"})
		return
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "plese provide correct user id password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": token})

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
