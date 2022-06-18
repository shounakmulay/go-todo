package controller

import (
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
)

type IRoleController interface {
	FindRoleByID(id int) (resmodel.Role, error)
	CreateRole(role reqmodel.CreateRole) (int, error)
}
