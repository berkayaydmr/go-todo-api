package repository

import (
	"errors"
	"go-todo-api/entities"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Insert(user *entities.User) (error)
	Update(user *entities.User) (error)
	Delete(user *entities.User) (error)
	FindAll(user []*entities.User) ([]*entities.User, error)
	FindByID(user *entities.User) (*entities.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (repository *UserRepository) Insert(user *entities.User) (error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) Update(user *entities.User) (error) {
	err := repository.db.Model(&user).Where("user_id = ?", user.Id).Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) Delete(user *entities.User) error {
	err := repository.db.Where("user_id = ?", user.Id).Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) FindAll(user []*entities.User) ([]*entities.User, error) {
	err := repository.db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (repository *UserRepository) FindByID(user *entities.User) (*entities.User, error) {
	err := repository.db.First(&user, int(user.Id)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil,nil
		}
		return nil, err
	}
	return user, nil
}