package helper

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"ecommerce.com/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashPassword(p string) (string, error) {

	if len(p) < 8 {
		return "", errors.New("password length should be at least 8 characters long")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), 10)

	if err != nil {
		log.Printf("Internal error on hash generation %v", err)
	}

	return string(hashedPassword), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {

	//  || role == ""
	if id == 0 || email == "" {
		return "", errors.New("required some valid inputs")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		return "", errors.New("unable to signed the token")
	}
	// ctx.SetSameSite(http.SameSiteLaxMode)
	// ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	return tokenString, nil
}

func (a Auth) VerifyPassword(pP string, hP string) error {

	if len(pP) < 8 {
		return errors.New("password length should be at least 8 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP))

	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}

func (a Auth) VerifyToken(t string) (domain.User, error) {
	tokenArr := strings.Split(t, " ")

	if len(tokenArr) != 2 || tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token format")
	}

	tokenStr := tokenArr[1]

	fmt.Printf("Token string: %s\n", tokenStr)

	// Parse the token string
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, fmt.Errorf("token parse error: %v", err)
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Access the claims and extract user information
		user := domain.User{
			ID:       uint(claims["user_id"].(float64)),
			Email:    claims["email"].(string),
			UserType: claims["role"].(string),
		}

		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorization(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		ctx.Abort()
		return
	}

	user, err := a.VerifyToken(authHeader)
	//jwt: token contains an invalid number of segments")
	log.Print("err123,", user, err)

	if err == nil && user.ID > 0 {
		ctx.Set("user", user)
		ctx.Next()

	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization failed"})
		return
	}

}
func (a Auth) GetCurrentUser(ctx *gin.Context) domain.User {
	user, exists := ctx.Get("user")

	if !exists {
		log.Fatalf("User doesn't exist in context")
	}

	

	if user, ok := user.(domain.User); ok {
		// Debug logging
	log.Printf("User data retrieved from context: %v", user)
	
	} else {
		log.Fatalf("User data is not of type domain.User")
	}

	// Return a default user or handle the error as needed
	return user.(domain.User) // Or handle the error in an appropriate way
}
func (a Auth) GenerateCode() (int, error) {
	return RandomCodeGeneration(6)
}
