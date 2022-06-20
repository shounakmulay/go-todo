package daos

import "go-todo/server/model/dbmodel"

type IUserDao interface {
	CreateUser(user dbmodel.User) (int, error)
	UpdateUser(user dbmodel.User) error
	UpdateUserToken(user *dbmodel.User) error
	FindUserByUsername(username string) (dbmodel.User, error)
	FindUserByID(id int) (dbmodel.User, error)
	DeleteUser(id int) error
}
