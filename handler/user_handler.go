package handler

import "go-todo-api/repository"

type UserHandler struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserHandler(UserRepo res)  {
	
}