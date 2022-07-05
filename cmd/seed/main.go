package main

import (
	"os"

	"go-todo/db/seed"
	"go-todo/internal/env"
	errorutl "go-todo/internal/error"
	"go-todo/internal/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	env.Load()

	loggerError := log.Logger.InitError
	errorutl.Fatal(loggerError)

	dsn := os.Getenv("DB_SQL_URL")

	safeEnvs := map[string]bool{
		"local":   true,
		"docker":  true,
		"develop": true,
	}

	environment := env.GetString("ENVIRONMENT_NAME")
	isSafe, ok := safeEnvs[environment]

	if !isSafe || !ok {
		log.Logger.Fatal("The command is not safe to run in this environment. " +
			"If you are sure you want to run it in this environment, you need to add it to the map of safe environments.")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errorutl.Fatal(err)

	seed.DB(db)
}
