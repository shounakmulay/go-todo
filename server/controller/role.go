package controller

import (
	"go-todo/internal/returns"
	"go-todo/server/daos"
	"go-todo/server/model/resmodel"
)

type RoleController struct {
	Dao *daos.RoleDao
}

func NewRoleController(roleDao *daos.RoleDao) *RoleController {
	return &RoleController{
		Dao: roleDao,
	}
}

func (r RoleController) GetRole(id int) (resmodel.Role, error) {
	role, err := r.Dao.FindRoleById(id)
	return returns.ErrorOrElse(err, func() resmodel.Role {
		return resmodel.Role{
			ID:          role.ID,
			Name:        role.Name,
			AccessLevel: role.AccessLevel,
		}
	})
}
