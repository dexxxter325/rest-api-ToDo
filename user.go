package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"` //binding:"required"-поле не может быть пустым
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
