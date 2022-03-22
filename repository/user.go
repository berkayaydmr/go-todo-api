package repository

import (
	"errors"
	"go-todo-api/entities"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Insert(user *entities.User) error
	Update(user *entities.User) error
	FindAll(user []*entities.User) ([]*entities.User, error)
	FindByID(user *entities.User) (*entities.User, error)
	FindByEmail(user *entities.User) (*entities.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (repository *UserRepository) Insert(user *entities.User) error {
	err := repository.db.Create(&user).Error
	return err
}

func (repository *UserRepository) Update(user *entities.User) error {
	err := repository.db.Model(&user).Save(&user).Error
	return err
}

func (repository *UserRepository) FindAll(user []*entities.User) ([]*entities.User, error) {
	err := repository.db.Not("status = ?", "PASSIVE").Find(&user).Error
	return user, err
}

func (repository *UserRepository) FindByID(user *entities.User) (*entities.User, error) {
	err := repository.db.Not("status = ?", "PASSIVE").Find(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) FindByEmail(user *entities.User) (*entities.User, error) {
	err := repository.db.Not("status = ?", "PASSIVE").Where("email = ?", user.Email).Find(&user).Error
	return user, err
}
