package daos

import (
	"go-todo/internal/returns"
	"go-todo/server/model/dbmodel"
	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *RoleDao {
	return &RoleDao{
		db: db,
	}
}

func (d RoleDao) FindRoleById(id int) (dbmodel.Role, error) {
	role := &dbmodel.Role{ID: id}
	result := d.db.First(role)
	return *role, result.Error
}

func (d RoleDao) CreateRole(role dbmodel.Role) (int, error) {
	result := d.db.Create(&role)
	return returns.ValueOrError(role.ID, result.Error)
}
