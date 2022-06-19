package controller

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"go-todo/server/config"
	"go-todo/server/model"
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/resmodel"
	"strconv"
	"time"
)

type JwtController struct {
	secret []byte
	ttl    time.Duration
	algo   jwt.SigningMethod
}

func NewJwtController(config *config.JWT) (JwtController, error) {
	secretLen := len(config.Secret)

	if secretLen < config.MinSecretLength {
		return JwtController{},
			errors.New(fmt.Sprintf("JWT secret length too short. Should be at least %v", config.MinSecretLength))
	}

	return JwtController{
		secret: []byte(config.Secret),
		ttl:    time.Duration(config.DurationMinutes) * time.Minute,
		algo:   jwt.GetSigningMethod(config.SigningAlgorithm),
	}, nil
}

func (c JwtController) GenerateTokens(user dbmodel.User) (resmodel.JwtTokens, error) {
	claims := &model.JwtCustomClaims{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.RoleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(c.ttl)),
		},
	}
	err := validator.New().Struct(claims)
	if err != nil {
		return resmodel.JwtTokens{}, nil
	}

	token, tokenErr := jwt.NewWithClaims(c.algo, claims).SignedString(c.secret)
	if tokenErr != nil {
		return resmodel.JwtTokens{}, err
	}

	sha := sha1.New()
	_, formatErr := fmt.Fprintf(sha, "%s%s", token, strconv.Itoa(time.Now().Nanosecond()))
	if formatErr != nil {
		return resmodel.JwtTokens{}, formatErr
	}
	refreshToken := fmt.Sprintf("%x", sha.Sum(nil))

	return resmodel.JwtTokens{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
