package controller

import (
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

	return r.Dao.CreateRole(dbRole)
}

func (r RoleController) FindRoleByID(id int) (resmodel.Role, error) {
	role, err := r.Dao.FindRoleByID(id)

	roleRes := resmodel.Role{
		ID:          role.ID,
		Name:        role.Name,
		AccessLevel: role.AccessLevel,
	}
	return roleRes, err
}
