package service

import (
	todo "REST_API_ToDo"
	"REST_API_ToDo/pkg/repository"
)

// service /*должен уметь общаться с бд*/
type Authorization interface {
	CreateUser(user todo.User) (int, error) //int=ID user`а
}
type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Authorization //авторизация
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
