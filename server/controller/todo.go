package controller

import (
	errutl "go-todo/internal/error"
	"go-todo/server/cache"
	"go-todo/server/daos"
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
)

type TodoController struct {
	dao   daos.ITodoDao
	cache cache.ITodoCache
}

func NewTodoController(todoDao daos.ITodoDao, todoCache cache.ITodoCache) *TodoController {
	return &TodoController{
		dao:   todoDao,
		cache: todoCache,
	}
}
func (t TodoController) CreateTodo(todo reqmodel.CreateTodo, userID int) (*resmodel.CreateTodo, error) {
	dbTodo := dbmodel.Todo{
		UserID:      userID,
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		Done:        todo.Done,
	}
	todoID, err := t.dao.CreateTodo(dbTodo)

	errutl.Log(t.cache.Invalidate(userID))

	return &resmodel.CreateTodo{
		ID: todoID,
	}, err
}

func (t TodoController) UpdateTodo(todo reqmodel.UpdateTodo, userID int) error {
	dbTodo := dbmodel.Todo{
		ID:          todo.ID,
		UserID:      userID,
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		Done:        todo.Done,
	}

	errutl.Log(t.cache.Invalidate(userID))

	return t.dao.UpdateTodo(dbTodo)
}

func (t TodoController) GetTodo(id int, userID int) (*resmodel.Todo, error) {
	dbTodo, err := t.dao.GetTodo(id, userID)
	if err != nil {
		return nil, err
	}

	resTodo := &resmodel.Todo{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: dbTodo.Description,
		DueDate:     dbTodo.DueDate,
		Done:        dbTodo.Done,
	}

	return resTodo, nil
}

func (t TodoController) GetAllTodos(userID int) (*[]resmodel.Todo, error) {
	cacheTodos, cacheErr := t.cache.GetAllTodos(userID)
	if cacheErr == nil {
		return &cacheTodos, nil
	}

	todos, err := t.dao.GetAllTodos(userID)
	if err != nil {
		return nil, err
	}

	resTodos := getResTodos(todos)
	errutl.Log(t.cache.SetAllTodos(userID, *resTodos))
	return resTodos, nil
}

func (t TodoController) GetAllTodosByState(done int8, userID int) (*[]resmodel.Todo, error) {
	todos, err := t.dao.GetAllTodosByState(done, userID)
	if err != nil {
		return nil, err
	}

	return getResTodos(todos), nil
}

func (t TodoController) DeleteTodo(id int, userID int) error {
	errutl.Log(t.cache.Invalidate(userID))
	return t.dao.DeleteTodo(id, userID)
}

func getResTodos(todos *[]dbmodel.Todo) *[]resmodel.Todo {
	var resTodos = []resmodel.Todo{}
	for _, todo := range *todos {
		resTodo := resmodel.Todo{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			DueDate:     todo.DueDate,
			Done:        todo.Done,
		}
		resTodos = append(resTodos, resTodo)
	}

	return &resTodos
}
