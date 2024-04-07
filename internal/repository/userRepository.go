package repository

import (
	"ecommerce.com/internal/domain"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type UserRepository interface {
	CreateUser(usr domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {

	return &userRepository{
		db: db,
	}

}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {

	err := r.db.Create(&usr).Error

	if err != nil {
		log.Printf("create error %v", err)
		return domain.User{}, errors.New("failed to create users")
	}
	return usr, nil

}

func (r userRepository) FindUser(email string) (domain.User, error) {

	var user domain.User

	err := r.db.First(&user, "email=?", email).Error

	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user doesnt exist")
	}

	return user, nil

}

func (r userRepository) FindUserById(id uint) (domain.User, error) {

	var user domain.User

	err := r.db.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("user with ID %d not found", id)
			return domain.User{}, errors.New("user not found")
		}
		// Handle other database errors
		log.Printf("find user error: %v", err)
		return domain.User{}, err

	}

	return user, nil

}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error

	if err != nil {
		log.Printf("update user error%v", err)
		return domain.User{}, errors.New("failed to update")
	}

	return user, nil

}
