package main

import (
	"fmt"
	"github.com/joho/godotenv"
	errorutl "go-todo/internal/error"
	"go-todo/internal/log"
	"go-todo/server"
	"go-todo/server/config"
	"os"
)

func main() {
	const envName = "ENVIRONMENT_NAME"
	env := os.Getenv(envName)

	if env == "" {
		env = "local"
		errorutl.Log(os.Setenv(envName, env))
	}

	errorutl.Fatal(
		godotenv.Load(fmt.Sprintf(".env.%s", env)),
	)

	loggerError := log.Logger.InitError
	errorutl.Fatal(loggerError)

	cfg, cfgErr := config.Load()
	errorutl.Fatal(cfgErr)

	err := server.Start(cfg)
	errorutl.Fatal(err)
}
