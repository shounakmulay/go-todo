package claims

import "github.com/golang-jwt/jwt/v4"

type JwtRefreshClaims struct {
	ID       int    `json:"ID" validate:"required"`
	Username string `json:"username" validate:"required"`
	jwt.RegisteredClaims
}
