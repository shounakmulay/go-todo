package daos

import (
	"go-todo/server/model/dbmodel"
)

type IRoleDao interface {
	CreateRole(role dbmodel.Role) (int, error)
	FindRoleByID(id int) (dbmodel.Role, error)
}
