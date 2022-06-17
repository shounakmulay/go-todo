package controller

import (
	"go-todo/server/model/resmodel"
)

type IRoleController interface {
	GetRole(id int) (resmodel.Role, error)
}
