package handler

import "go-todo-api/repository"

type UserHandler struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserHandler(UserRepo repository.UserRepositoryInterface) *UserHandler {
	return &UserHandler{UserRepository: UserRepo}
}

