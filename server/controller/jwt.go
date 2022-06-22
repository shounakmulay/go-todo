package controller

import (
	"fmt"
	"time"

	"go-todo/server/config"
	"go-todo/server/model/claims"
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/resmodel"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
)

type JwtController struct {
	secret        []byte
	refreshSecret []byte
	ttl           time.Duration
	refreshTTL    time.Duration
	algo          jwt.SigningMethod
}

func NewJwtController(config *config.JWT) (JwtController, error) {
	secretLen := len(config.Secret)

	if secretLen < config.MinSecretLength {
		return JwtController{},
			fmt.Errorf("JWT secret length too short. Should be at least %v", config.MinSecretLength)
	}

	return JwtController{
		secret:        []byte(config.Secret),
		refreshSecret: []byte(config.RefreshSecret),
		ttl:           time.Duration(config.DurationMinutes) * time.Minute,
		refreshTTL:    time.Duration(config.RefreshDurationMinutes) * time.Minute,
		algo:          jwt.GetSigningMethod(config.SigningAlgorithm),
	}, nil
}

func (c JwtController) GenerateTokens(user dbmodel.User) (resmodel.JwtTokens, error) {
	jwtclaims := &claims.JwtClaims{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(c.ttl)),
		},
	}
	refreshClaims := &claims.JwtRefreshClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(c.refreshTTL)),
			Subject:   "refresh",
		},
	}
	err := validator.New().Struct(jwtclaims)
	if err != nil {
		return resmodel.JwtTokens{}, nil
	}

	token, tokenErr := jwt.NewWithClaims(c.algo, jwtclaims).SignedString(c.secret)
	if tokenErr != nil {
		return resmodel.JwtTokens{}, tokenErr
	}

	refreshToken, refreshTokenErr := jwt.NewWithClaims(c.algo, refreshClaims).SignedString(c.refreshSecret)
	if refreshTokenErr != nil {
		return resmodel.JwtTokens{}, refreshTokenErr
	}

	return resmodel.JwtTokens{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (c JwtController) GetRefreshSecret() []byte {
	return c.refreshSecret
}
