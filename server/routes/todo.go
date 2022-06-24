package routes

import (
	errorutl "go-todo/internal/error"
	"go-todo/internal/json"
	"go-todo/server/controller"
	"go-todo/server/model/ctx"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
	"go-todo/server/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Todo(g *echo.Group, controller controller.ITodoController) {
	todo := g.Group("/todo")

	todo.POST("", func(c echo.Context) error {
		user := ctx.GetUserFromContext(c)
		reqTodo := &reqmodel.CreateTodo{}

		bindValErr := validator.BindAndValidateWith(c, reqTodo, validator.BindBody)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		resTodo, err := controller.CreateTodo(*reqTodo, user.ID)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, user.ID))
		}

		return json.Created(c, resTodo)
	})

	todo.PUT("", func(c echo.Context) error {
		user := ctx.GetUserFromContext(c)
		reqTodo := &reqmodel.UpdateTodo{}

		bindValErr := validator.BindAndValidateWith(c, reqTodo, validator.BindBody)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		err := controller.UpdateTodo(*reqTodo, user.ID)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, user.ID))
		}

		return c.NoContent(http.StatusOK)
	})

	todo.DELETE("/:id", func(c echo.Context) error {
		user := ctx.GetUserFromContext(c)
		todoID := &reqmodel.TodoID{}

		bindValErr := validator.BindAndValidateWith(c, todoID, validator.BindPath)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		err := controller.DeleteTodo(todoID.ID, user.ID)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, user.ID))
		}

		return c.NoContent(http.StatusOK)
	})

	todo.GET("/:id", func(c echo.Context) error {
		user := ctx.GetUserFromContext(c)
		todoID := &reqmodel.TodoID{}

		bindValErr := validator.BindAndValidateWith(c, todoID, validator.BindPath)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		todo, err := controller.GetTodo(todoID.ID, user.ID)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, todoID.ID))
		}

		return json.Success(c, todo)
	})

	todo.GET("/list", func(c echo.Context) error {
		user := ctx.GetUserFromContext(c)
		var todoStatus *reqmodel.TodoStatus
		var todoResults *[]resmodel.Todo

		if c.QueryParam("done") != "" {
			// Has query param to filter by state

			todoStatus = &reqmodel.TodoStatus{}

			bindValErr := validator.BindAndValidateWith(c, todoStatus, validator.BindQuery)
			if bindValErr != nil {
				return json.Error(c, bindValErr)
			}

			todos, err := controller.GetAllTodosByState(int8(todoStatus.Done), user.ID)
			if err != nil {
				return json.Error(c, errorutl.GormToResErr(err, user.ID))
			}
			todoResults = todos
		} else {
			// Does not have query param to filter by state

			todos, err := controller.GetAllTodos(user.ID)
			if err != nil {
				return json.Error(c, errorutl.GormToResErr(err, user.ID))
			}
			todoResults = todos
		}

		return json.Success(c, todoResults)
	})
}
