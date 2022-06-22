package controller

import (
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
)

type ITodoController interface {
	CreateTodo(todo reqmodel.CreateTodo, userID int) (*resmodel.CreateTodo, error)
	UpdateTodo(todo reqmodel.UpdateTodo, userID int) error
	GetTodo(id int, userID int) (*resmodel.Todo, error)
	GetAllTodos(userID int) (*[]resmodel.Todo, error)
	GetAllTodosByState(done int8, userID int) (*[]resmodel.Todo, error)
	DeleteTodo(id, userID int) error
}
