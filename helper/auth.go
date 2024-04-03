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

	if id == 0 || email == "" || role == "" {
		return "", errors.New("required some valid inputs")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		return "", errors.New("unable to signed the token")
	}

	return tokenString, nil
}

func (a Auth) VerifyPassword(hP string, pP string) error {

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

	tokenStr := tokenArr[0]
	if tokenStr[:7] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	// jwt.verify
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)

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
	if err != nil || user.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}

func (a Auth) GetCurrentUser(ctx *gin.Context) domain.User {

	user := ctx.MustGet("user")

	return user.(domain.User)

}
