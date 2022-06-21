package controller

import (
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/resmodel"
)

type IJwtController interface {
	GenerateTokens(user dbmodel.User) (resmodel.JwtTokens, error)
	GetRefreshSecret() []byte
}
