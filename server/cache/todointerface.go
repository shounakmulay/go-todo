package cache

import "go-todo/server/model/resmodel"

type ITodoCache interface {
	GetAllTodos(userID int) ([]resmodel.Todo, error)
	SetAllTodos(userID int, todos []resmodel.Todo) error
	Invalidate(userID int) error
}
