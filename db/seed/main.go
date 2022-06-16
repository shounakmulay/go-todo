package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
	errorutl "go-todo/internal/util/error"
	"go-todo/internal/util/log"
	"go-todo/model/dbmodel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"time"
)

func main() {
	log.InitLogger()

	errorutl.Panic(godotenv.Load(".env.develop"))
	dsn := os.Getenv("SQL_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errorutl.Fatal(err)

	seedRoles(db)
	seedUsers(db)
	seedTodos(db)
}

func seedRoles(db *gorm.DB) {
	log.Logger.Debug("Seeding Roles...")
	roles := []dbmodel.Role{
		{
			Name:        "user",
			AccessLevel: 100,
		},
		{
			Name:        "auditor",
			AccessLevel: 200,
		},
		{
			Name:        "admin",
			AccessLevel: 300,
		},
	}
	errorutl.Fatal(db.Create(&roles).Error)
	log.Logger.Debug("Seeding Roles Done.")
}

func seedUsers(db *gorm.DB) {
	log.Logger.Debug("Seeding Users...")
	var roles []dbmodel.Role
	result := db.Find(&roles)
	errorutl.Fatal(result.Error)

	var users []dbmodel.User
	// add users
	for i := 0; i < 8; i++ {
		user := dbmodel.User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Username:  gofakeit.Username(),
			Password:  gofakeit.Password(true, true, true, true, false, 16),
			Email:     gofakeit.Email(),
			Mobile:    gofakeit.Phone(),
			RoleId:    roles[0].ID,
		}
		users = append(users, user)
	}
	// add auditors
	for i := 0; i < 4; i++ {
		user := dbmodel.User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Username:  gofakeit.Username(),
			Password:  gofakeit.Password(true, true, true, true, false, 16),
			Email:     gofakeit.Email(),
			Mobile:    gofakeit.Phone(),
			RoleId:    roles[1].ID,
		}
		users = append(users, user)
	}
	// add admins
	for i := 0; i < 3; i++ {
		user := dbmodel.User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Username:  gofakeit.Username(),
			Password:  gofakeit.Password(true, true, true, true, false, 16),
			Email:     gofakeit.Email(),
			Mobile:    gofakeit.Phone(),
			RoleId:    roles[2].ID,
		}
		users = append(users, user)
	}
	errorutl.Fatal(db.Create(&users).Error)
	log.Logger.Debug("Seeding Users Done.")
}

func seedTodos(db *gorm.DB) {
	log.Logger.Debug("Seeding Todos...")
	var users []dbmodel.User
	result := db.Where(&dbmodel.User{RoleId: 1}).Find(&users)
	errorutl.Fatal(result.Error)

	var todos []dbmodel.Todo
	for _, u := range users {
		r := rand.Intn(10)
		for i := 0; i <= r; i++ {
			todo := dbmodel.Todo{
				UserID:      u.ID,
				Title:       gofakeit.Sentence(10),
				Description: gofakeit.Sentence(100),
				DueDate:     gofakeit.DateRange(time.Unix(0, 0), time.Now()),
			}
			todos = append(todos, todo)
		}
	}
	errorutl.Fatal(db.Create(&todos).Error)
	log.Logger.Debug("Seeding Todos Done.")
}
