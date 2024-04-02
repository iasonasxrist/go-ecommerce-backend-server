package service

import (
	"log"

	"ecommerce.com/internal/domain"
	"ecommerce.com/internal/dto"
	"ecommerce.com/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) FindByEmail(email string) (*domain.User, error) {

	user, err := s.Repo.FindUser(email)
	return &user, err
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {

	user, err := s.Repo.CreateUser(domain.User{
		Email: input.Email,
		Password: input.Password,
		Phone: input.Phone,
	})
	return "this-is-my-token", nil
}

func (s *UserService) Login(input any) (string, error) {
	// Perform database operations to log in the user
	// For example:
	// sessionToken, err := database.LoginUser(email, password)
	// if err != nil {
	//     return "", err
	// }
	// return sessionToken, nil
	return "", nil // Placeholder return
}

func (s *UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil
}

func (s *UserService) VerifyCode(id uint, code int) error {

	return nil
}

func (s *UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (s *UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil
}

func (s *UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s *UserService) BecomeSeller(id uint, input any) (string, error) {

	return "", nil
}

func (s *UserService) FindCart(id uint) ([]interface{}, error) {

	var result []interface{}
	return result, nil
}

func (s *UserService) CreateCart(id uint, u domain.User) ([]interface{}, error) {

	var result []interface{}
	return result, nil
}

func (s *UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s *UserService) GetOrders(u domain.User) ([]interface{}, error) {
	var result []interface{}
	return result, nil
}

func (s *UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	var result []interface{}
	return result, nil
}
