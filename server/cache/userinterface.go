package cache

import "go-todo/server/model/resmodel"

type IUserCache interface {
	GetUserByUsername(username string) (*resmodel.User, error)
	GetUserByID(id int) (*resmodel.User, error)
	SetUserForUsername(username string, user resmodel.User) error
	SetUserForID(id int, user resmodel.User) error
	Invalidate(id int, username string) error
}
