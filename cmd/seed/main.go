package main

import (
	"github.com/joho/godotenv"
	"go-todo/db/seed"
	errorutl "go-todo/internal/error"
	"go-todo/internal/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
)

func main() {
	loggerError := log.Logger.InitError
	errorutl.Fatal(loggerError)

	errorutl.Panic(godotenv.Load(".env.local"))
	dsn := os.Getenv("DB_SQL_URL")

	isLocal := strings.Contains(dsn, "localhost") || strings.Contains(dsn, "127.0.0.1")
	if !isLocal {
		log.Logger.Fatal("The command is not running on a local db instance. " +
			"If you are sure you want to run it on a non local instance, you need to disable this check manually")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errorutl.Fatal(err)

	seed.DB(db)
}
