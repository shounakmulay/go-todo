package claims

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	ID       int    `json:"ID" validate:"required"`
	Username string `json:"username" validate:"required"`
	Role     int    `json:"role" validate:"required"`
	jwt.RegisteredClaims
}
