package service

import (
	todo "REST_API_ToDo"
	"REST_API_ToDo/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "siakgj7435ty438jmsd"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService { //конструктор для работы с бд
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user todo.User) (int, error) { //новый пользователь в бд
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func generatePasswordHash(password string) string {
	hash := sha1.New()                               //хэш пароля
	hash.Write([]byte(password))                     //байтовый пароль
	return fmt.Sprintf("%x", hash.Sum([]byte(salt))) //hash.Sum-итоговый хэш пароля в виде байтового массива.%x-шестнадцатеричныя запись
}
