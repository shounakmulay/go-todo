package controller

import (
	"go-todo/internal/returns"
	"go-todo/server/daos"
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
)

type RoleController struct {
	Dao daos.IRoleDao
}

func NewRoleController(roleDao daos.IRoleDao) *RoleController {
	return &RoleController{
		Dao: roleDao,
	}
}

func (r RoleController) CreateRole(role reqmodel.CreateRole) (int, error) {
	dbRole := dbmodel.Role{
		Name:        role.Name,
		AccessLevel: role.AccessLevel,
	}
	id, err := r.Dao.CreateRole(dbRole)

	return returns.ErrorOrValue(err, id)
}

func (r RoleController) FindRoleByID(id int) (resmodel.Role, error) {
	role, err := r.Dao.FindRoleById(id)

	return returns.ErrorOrValue(err, resmodel.Role{
		ID:          role.ID,
		Name:        role.Name,
		AccessLevel: role.AccessLevel,
	})
}
