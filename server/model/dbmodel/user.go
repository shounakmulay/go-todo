package dbmodel

import (
	"errors"
	"time"

	"go-todo/internal/log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
	Mobile    string
	Token     string
	RoleID    int
	Role      Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		errorMsg := "error hashing password"
		log.Logger.Errorf(errorMsg)

		return errors.New(errorMsg)
	}

	u.Password = string(hashedPass)

	return nil
}
