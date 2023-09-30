CREATE TABLE users
/*unique-знач.внутри таблицы не будут повторяться
  NOT NULL-поле не может быть пустым
  serial- уникальное числовое значение*/
(

    id serial not null unique ,
    name varchar(255) not null ,
    username varchar(255) NOT NULL unique ,
    password_hash varchar(255) NOT NULL

);
CREATE TABLE todo_lists/*сам список задач */
(
    id serial not null unique ,
    title varchar(255) not null ,
    description varchar(255)
);
CREATE TABLE user_lists (/*списки пользователей*/
                            id serial NOT NULL UNIQUE,
                            user_id  int references users(id) on delete cascade NOT NULL,
                            list_id  int references todo_lists(id) on delete cascade NOT NULL
    /*references users(id)-внешний ключ(из 2х табл.ищет одинак.заголовки),связанные с todo_lists.in delete cascade-при удалении строки из таблицы users-соотв. строки в "user_lists" также будут удалены*/

);
CREATE TABLE todo_items (/*сами задачи в тодо*/
                            id serial NOT NULL UNIQUE,
                            title varchar(255) NOT NULL,
                            description varchar(255),
                            done BOOLEAN NOT NULL default false

);
CREATE TABLE lists_items (/*связь между списком и элементом в списке*/
                             id serial NOT NULL UNIQUE,
                             item_id int references todo_items (id) on delete cascade NOT NULL,
                             list_id int references todo_lists (id) on delete cascade NOT NULL



);