package controller

import (
	"go-todo/server/daos"
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
)

type UserController struct {
	dao daos.IUserDao
}

func NewUserController(dao daos.IUserDao) *UserController {
	return &UserController{dao: dao}
}

func (u UserController) CreateUser(user reqmodel.CreateUser) (int, error) {
	dbUser := dbmodel.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Mobile:    user.Mobile,
		RoleId:    user.RoleID,
	}

	return u.dao.CreateUser(dbUser)
}

func (u UserController) UpdateUser(user *dbmodel.User) error {
	return u.dao.UpdateUserToken(user)
}

func (u UserController) FindUserByUsername(username string) (resmodel.User, error) {
	dbUser, err := u.dao.FindUserByUsername(username)
	if err != nil {
		return resmodel.User{}, err
	}

	return dbToResUser(dbUser), nil
}

func (u UserController) FindDBUserByUsername(username string) (dbmodel.User, error) {
	dbUser, err := u.dao.FindUserByUsername(username)
	if err != nil {
		return dbmodel.User{}, err
	}

	return dbUser, nil
}

func (u UserController) FindUser(id int) (resmodel.User, error) {
	dbUser, err := u.dao.FindUserByID(id)
	if err != nil {
		return resmodel.User{}, err
	}

	return dbToResUser(dbUser), nil
}

func dbToResUser(dbUser dbmodel.User) resmodel.User {
	return resmodel.User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Mobile:    dbUser.Mobile,
		RoleID:    dbUser.RoleId,
	}
}
