package main

import (
	"go-todo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/go_todos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	roles := []model.Role{
		{
			Name:        "role1",
			AccessLevel: 1,
		},
		{
			Name:        "role2",
			AccessLevel: 2,
		},
		{
			Name:        "role3",
			AccessLevel: 3,
		},
	}
	err = db.Create(&roles).Error
	if err != nil {
		log.Fatalln(err)
	}
}
