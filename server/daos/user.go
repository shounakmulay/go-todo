package daos

import (
	"go-todo/server/model/dbmodel"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (u UserDao) CreateUser(user dbmodel.User) (int, error) {
	result := u.db.Create(&user)
	return user.ID, result.Error
}

func (u UserDao) UpdateUser(user *dbmodel.User) error {
	result := u.db.Save(&user)
	return result.Error
}

func (u UserDao) FindUserByUsername(username string) (dbmodel.User, error) {
	user := &dbmodel.User{
		Username: username,
	}
	result := u.db.Where("username = ?", user.Username).First(user)

	return *user, result.Error
}

func (u UserDao) FindUserByID(id int) (dbmodel.User, error) {
	user := &dbmodel.User{}
	result := u.db.First(user, id)

	return *user, result.Error
}
