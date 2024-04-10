package services

import (
	"fmt"
	"net/http"
	"strconv"

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
		Repo:   repository.NewRepository(rh.Db),
		Auth:   rh.Auth,
		Config: rh.Config,
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
	fmt.Printf("User code: %s\n", user.Code) // Print user code for debugging purposes

	var req map[string]interface{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please provide a valid input"})
		return
	}

	codeStr, ok := req["code"].(string)
	fmt.Print("CODEsTR", codeStr)
	if !ok {
		fmt.Println("Invalid code type") // Print error message for debugging purposes
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid code type"})
		return
	}

	code, err := strconv.Atoi(codeStr)
	if err != nil {
		fmt.Println("Error converting code to int:", err) // Print error message for debugging purposes
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error converting code to int"})
		return
	}

	fmt.Printf("Request code:  %v, %d\n", user.Code, code) // Print request code for debugging purposes

	// Check if user code matches request code
	if user.Code != code {
		fmt.Println("User code does not match request code") // Print error message for debugging purposes
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "verification code does not match"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "verified successfully"})
}
func (h *UserHandler) GetVerificationCode(ctx *gin.Context) {

	user := h.svc.Auth.GetCurrentUser(ctx)

	// replaced two returns by only one
	err := h.svc.GetVerificationCode(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to generate verification code"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "verification Code"})

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
