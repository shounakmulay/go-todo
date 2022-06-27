package daos

import (
	errorutl "go-todo/internal/error"
	"go-todo/server/model/dbmodel"

	"gorm.io/gorm"
)

type TodoDao struct {
	db *gorm.DB
}

func NewTodoDao(db *gorm.DB) *TodoDao {
	return &TodoDao{
		db: db,
	}
}

func (t TodoDao) CreateTodo(todo dbmodel.Todo) (int, error) {
	result := t.db.Create(&todo)
	return todo.ID, result.Error
}

func (t TodoDao) UpdateTodo(todo dbmodel.Todo) error {
	result := t.db.Model(todo).Where(dbmodel.Todo{
		ID:     todo.ID,
		UserID: todo.UserID,
	}).Updates(&todo)

	if result.RowsAffected == 0 {
		return errorutl.NewQueryError("No matching data found to update")
	}

	return result.Error
}

func (t TodoDao) GetTodo(id, userID int) (*dbmodel.Todo, error) {
	todo := &dbmodel.Todo{
		UserID: userID,
	}
	result := t.db.Where(todo).First(todo, id)
	return todo, result.Error
}

func (t TodoDao) GetAllTodos(userID int) (*[]dbmodel.Todo, error) {
	todos := &[]dbmodel.Todo{}
	result := t.db.Where(&dbmodel.Todo{UserID: userID}).Find(todos)
	return todos, result.Error
}

func (t TodoDao) GetAllTodosByState(done int8, userID int) (*[]dbmodel.Todo, error) {
	if done != 1 && done != 0 {
		return nil, errorutl.NewQueryError(`invalid value for "done". Should be either 0 or 1`)
	}
	todos := &[]dbmodel.Todo{}
	result := t.db.Where(&dbmodel.Todo{Done: done, UserID: userID}).Find(todos)
	return todos, result.Error
}

func (t TodoDao) DeleteTodo(id int, userID int) error {
	dbUser := dbmodel.Todo{
		ID:     id,
		UserID: userID,
	}
	result := t.db.Where(&dbmodel.Todo{UserID: userID}).Delete(&dbUser)

	if result.RowsAffected == 0 {
		return errorutl.NewQueryError("No data to delete")
	}

	return result.Error
}
