package daos

import "go-todo/server/model/dbmodel"

type ITodoDao interface {
	CreateTodo(todo dbmodel.Todo) (int, error)
	UpdateTodo(todo dbmodel.Todo) error
	GetTodo(id, userID int) (*dbmodel.Todo, error)
	GetAllTodos(userID int) (*[]dbmodel.Todo, error)
	GetAllTodosByState(done int8, userID int) (*[]dbmodel.Todo, error)
	DeleteTodo(id, userID int) error
}
