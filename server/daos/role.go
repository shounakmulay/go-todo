package daos

import (
	"go-todo/server/model/dbmodel"

	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{
		db: db,
	}
}

func (d RoleDao) FindRoleByID(id int) (dbmodel.Role, error) {
	role := &dbmodel.Role{ID: id}
	result := d.db.First(role)
	return *role, result.Error
}

func (d RoleDao) CreateRole(role dbmodel.Role) (int, error) {
	result := d.db.Create(&role)
	return role.ID, result.Error
}
