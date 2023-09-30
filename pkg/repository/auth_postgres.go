package repository

import (
	todo "REST_API_ToDo"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB //соед. с бд
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user todo.User) (int, error) { //созд.нового пользователя
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable) //INSERT INTO-вставить строку в нашу таблу
	//RETURNING id будет возвращать id новой записи
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password) //QueryRow выполняет sql запрос и возвращает ток 1 строку,query-строка с самим SQL запросом,остальное подставляется в цифры сверху(плейсхолдеры)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
