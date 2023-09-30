package repository

import (
	todo "REST_API_ToDo"
	"github.com/jmoiron/sqlx"
)

/*взаимодействие с бд*/

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}
type TodoList interface {
}
type TodoItem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
