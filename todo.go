package todo

type Todolist struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type UsersList struct { //списки пользователей
	Id     int //индефицирует каждый эл в спике
	UserId int
	ListId int //индефикатор списка
}
type TodoItem struct { //задача в тодо
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"` //выполнена ли задача
}
type ListsItem struct { //связь между списком и элементом в списке
	Id     int //уникальная связь между списком и эл.
	ListId int //список,кот. принадлежит текущий элемент
	ItemId int //элемент списка, с которым связана текущая связь
}
