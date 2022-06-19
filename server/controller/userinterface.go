package controller

import (
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
)

type IUserController interface {
	CreateUser(user reqmodel.CreateUser) (int, error)
	UpdateUser(user *dbmodel.User) error
	FindUserByUsername(username string) (resmodel.User, error)
	FindDBUserByUsername(username string) (dbmodel.User, error)
	FindUser(id int) (resmodel.User, error)
}
