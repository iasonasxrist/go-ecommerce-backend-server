package service

import (
	"errors"
	"fmt"
	"log"
	// "strconv"
	"time"

	"ecommerce.com/config"
	"ecommerce.com/helper"
	"ecommerce.com/internal/domain"
	"ecommerce.com/internal/dto"
	"ecommerce.com/internal/repository"
	"ecommerce.com/pkg/notification"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s UserService) FindByEmail(email string) (*domain.User, error) {

	user, err := s.Repo.FindUser(email)
	return &user, err
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {

	hPassword, err := s.Auth.CreateHashPassword(input.Password)

	if err != nil {
		return "", nil
	}
	user, _ := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})

	log.Printf("user created %v", user)

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) Login(email string, password string) (string, error) {

	user, err := s.FindByEmail(email)

	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}
	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}
	// generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) IsVerified(id uint) bool {

	currentUser, err := s.Repo.FindUserById(id)

	return err == nil && currentUser.Verified
}

func (s UserService) GetVerificationCode(e domain.User) (int,error) {

	if s.IsVerified(e.ID) {
		return 0, nil
	}

	code, err := s.Auth.GenerateCode()

	if err != nil {
		return 0, err
	}

	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repo.UpdateUser(e.ID, user)

	if err != nil {
		return 0,errors.New("unable to update verification code")
	}
	// user, err = s.Repo.FindUserById(user.ID)

	// Send SMS
	// notificationClient := notification.NewNotificationClient(s.Config)
	// notificationClient.SendSMS(user.Phone, strconv.Itoa(code))
	return code, nil
}

func (s UserService) VerifyCode(id uint, code int) error {

	if s.IsVerified(id) {
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserById(id)

	fmt.Printf("**** Correct user data ***** %v\n", user)
	// correct

	if err != nil {
		return errors.New(err.Error())
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	userUpdated := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, userUpdated)

	if err != nil {
		return errors.New("unable to verify unset")
	}

	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {

	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {

	var result []interface{}
	return result, nil
}

func (s UserService) CreateCart(id uint, u domain.User) ([]interface{}, error) {

	var result []interface{}
	return result, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	var result []interface{}
	return result, nil
}

func (s UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	var result []interface{}
	return result, nil
}
