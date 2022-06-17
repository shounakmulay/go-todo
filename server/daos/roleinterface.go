package daos

import (
	"go-todo/server/model/dbmodel"
)

type IRoleDao interface {
	CreateRole(role dbmodel.Role) (int, error)
	FindRoleById(id int) (dbmodel.Role, error)
}
