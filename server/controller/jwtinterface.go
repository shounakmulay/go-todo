package controller

import "go-todo/server/model/dbmodel"

type IJwtController interface {
	GenerateToken(user dbmodel.User)
}
