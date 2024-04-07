package services

import (
	"fmt"
	"log"
	"net/http"
	"ecommerce.com/helper"
	"ecommerce.com/internal/api/rest"
	"ecommerce.com/internal/dto"
	"ecommerce.com/internal/repository"
	"ecommerce.com/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc  service.UserService
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

	}

	token, err := h.svc.Signup(user)
	fmt.Printf("token 1%v", err)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})

	}

	ctx.JSON(http.StatusOK, gin.H{"message": token})

}

func (h *UserHandler) Login(ctx *gin.Context) {

	loginInput := dto.UserLogin{}

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

	user := h.svc.Auth.GetCurrentUser(ctx)
	fmt.Printf("*****user that bugs *******\n %v", user)
	var req dto.VerificationCodeInput

	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("code123 %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please provide a valid input"})
	}

	fmt.Printf("Type of user.ID: %T\n", user.ID)
	fmt.Printf("Type of req.Code: %T\n", req.Code)

	err := h.svc.VerifyCode(user.ID, req.Code)

	if err != nil {
		log.Printf("Error : %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "verified successfully"})

}

func (h *UserHandler) GetVerificationCode(ctx *gin.Context) {

	user := h.svc.Auth.GetCurrentUser(ctx)
	// fmt.Printf("userfgf %v", user)

	code, err := h.svc.GetVerificationCode(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to generate verification code"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "verification Code",
		"data": code})

}

func (h *UserHandler) GetProfile(ctx *gin.Context) {

	user := h.svc.Auth.GetCurrentUser(ctx)

	//TODO// Error on user set

	ctx.JSON(http.StatusOK, gin.H{"message": "Get Profile", "user": user})
}

func (h *UserHandler) CreateProfile(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"message": "Create Profile"})
}

func (h *UserHandler) GetCart(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"message": "Get Cart"})
}

func (h *UserHandler) AddToCart(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"message": "Add To Cart"})
}

func (h *UserHandler) GetOrders(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"message": "Get Orders"})
}

func (h *UserHandler) GetOrder(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"message": "Get Order"})
}

func (h *UserHandler) BecomeSeller(ctx *gin.Context) {

	// if err != nil {
	// 	logger.Error("connection error :%v", err)
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"message": "Become a Seller"})

}
