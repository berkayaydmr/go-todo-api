package repository

import (
	"errors"
	"go-todo-api/entities"

	"gorm.io/gorm"
)

type ToDoRepositoryInterface interface {
	Insert(i *entities.ToDo) error
	Update(i *entities.ToDo) (*entities.ToDo, error)
	Delete(i *entities.ToDo) error
	FindAll(result []*entities.ToDo) ([]*entities.ToDo, error)
	FindByID(result *entities.ToDo) (*entities.ToDo, error)
}

type ToDoRepository struct {
	db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) *ToDoRepository {
	return &ToDoRepository{db}
}

func (repository *ToDoRepository) Insert(i *entities.ToDo) (error) {
	err := repository.db.Create(&i).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *ToDoRepository) Update(i *entities.ToDo) (*entities.ToDo, error) {
	err := repository.db.Model(&i).Where("id = ?", i.Id).Save(&i).Error
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (repository *ToDoRepository) Delete(i *entities.ToDo) error {
	return repository.db.Delete(&i).Error
}

func (repository *ToDoRepository) FindAll(result []*entities.ToDo) ([]*entities.ToDo, error) {
	err := repository.db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (repository *ToDoRepository) FindByID(result *entities.ToDo) (*entities.ToDo, error) {
	err := repository.db.First(&result, int(result.Id)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil,nil
		}
		return nil, err
	}
	return result, nil
}
