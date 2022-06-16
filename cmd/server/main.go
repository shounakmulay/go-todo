package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-todo/server"
	"go-todo/server/config"
	errorutl "go-todo/utl/error"
	"go-todo/utl/log"
	"os"
)

func main() {
	log.InitLogger()
	const envName = "ENVIRONMENT_NAME"
	env := os.Getenv(envName)

	if env == "" {
		env = "local"
		errorutl.Log(os.Setenv(envName, env))
	}

	errorutl.Fatal(
		godotenv.Load(fmt.Sprintf(".env.%s", env)),
	)

	cfg, cfgErr := config.Load()
	errorutl.Fatal(cfgErr)

	server.Start(cfg)
}
