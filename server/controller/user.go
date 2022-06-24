package controller

import (
	errutl "go-todo/internal/error"
	"go-todo/server/cache"
	"go-todo/server/daos"
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
)

type UserController struct {
	dao   daos.IUserDao
	cache cache.IUserCache
}

func NewUserController(dao daos.IUserDao, cache cache.IUserCache) *UserController {
	return &UserController{dao: dao, cache: cache}
}

func (u UserController) CreateUser(user reqmodel.CreateUser) (int, error) {
	dbUser := dbmodel.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Mobile:    user.Mobile,
		RoleID:    user.RoleID,
	}

	return u.dao.CreateUser(dbUser)
}

func (u UserController) UpdateUser(user reqmodel.UpdateUser) error {
	dbUser := dbmodel.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Mobile:    user.Mobile,
	}
	errutl.Log(u.cache.Invalidate(dbUser.ID, dbUser.Username))
	return u.dao.UpdateUser(dbUser)
}

func (u UserController) UpdateUserToken(user *dbmodel.User) error {
	return u.dao.UpdateUserToken(user)
}

func (u UserController) FindUserByUsername(username string) (resmodel.User, error) {
	cacheUser, cacheErr := u.cache.GetUserByUsername(username)
	if cacheErr == nil {
		return *cacheUser, nil
	}

	dbUser, err := u.dao.FindUserByUsername(username)
	if err != nil {
		return resmodel.User{}, err
	}

	resUser := dbToResUser(dbUser)
	errutl.Log(u.cache.SetUserForUsername(username, resUser))
	return resUser, nil
}

func (u UserController) FindDBUserByUsername(username string) (dbmodel.User, error) {
	dbUser, err := u.dao.FindUserByUsername(username)
	if err != nil {
		return dbmodel.User{}, err
	}

	return dbUser, nil
}

func (u UserController) FindUser(id int) (resmodel.User, error) {
	cacheUser, cacheErr := u.cache.GetUserByID(id)
	if cacheErr == nil {
		return *cacheUser, nil
	}

	dbUser, err := u.dao.FindUserByID(id)
	if err != nil {
		return resmodel.User{}, err
	}

	resUser := dbToResUser(dbUser)
	errutl.Log(u.cache.SetUserForID(id, resUser))
	return resUser, nil
}

func (u UserController) DeleteUser(id int) error {
	return u.dao.DeleteUser(id)
}

func dbToResUser(dbUser dbmodel.User) resmodel.User {
	return resmodel.User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Mobile:    dbUser.Mobile,
		RoleID:    dbUser.RoleID,
	}
}
