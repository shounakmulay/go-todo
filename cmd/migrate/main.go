package main

import (
	"go-todo/internal/env"
	errorutl "go-todo/internal/error"
	"go-todo/internal/log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	env.Load()

	loggerError := log.Logger.InitError
	errorutl.Fatal(loggerError)

	dsn := os.Getenv("DB_SQL_URL")

	log.Logger.Info("Starting migration...")
	m, err := migrate.New("file:///go/src/go-todo/db/migration", "mysql://"+dsn)
	errorutl.Fatal(err)
	errorutl.Log(m.Up()) // or m.Step(2) if you want to explicitly set the number of migrations to run
	log.Logger.Info("Migration completed.")
}
