package main

import (
	"go-todo/internal/env"
	errorutl "go-todo/internal/error"
	"go-todo/internal/log"
	"go-todo/server"
	"go-todo/server/config"
)

func main() {
	env.Load()

	loggerError := log.Logger.InitError
	errorutl.Fatal(loggerError)

	cfg, cfgErr := config.Load()
	errorutl.Fatal(cfgErr)

	err := server.Start(cfg)
	errorutl.Fatal(err)
}
