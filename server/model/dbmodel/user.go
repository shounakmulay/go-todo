package dbmodel

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
	Mobile    string
	RoleId    int
	Role      Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
