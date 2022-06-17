package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-todo/db/mysql"
	"go-todo/internal/env"
	errorutl "go-todo/internal/error"
	"go-todo/internal/log"
	"go-todo/server/api"
	"go-todo/server/config"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start(cfg *config.Configuration) error {
	db, err := mysql.Connect(cfg)
	if err != nil {
		return errorutl.Format("error connecting to MySQL", err)
	}

	echoServer, echoError := api.Start(cfg, db)
	if echoError != nil {
		return errorutl.Format("error initializing api", echoError)
	}

	return startServer(cfg, echoServer)
}

func startServer(cfg *config.Configuration, echoServer *echo.Echo) error {

	httpServer := &http.Server{
		Addr:         cfg.Server.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeoutSeconds) * time.Second,
	}

	echoServer.Debug = cfg.Server.Debug

	// Start server
	go func() {
		log.Logger.Infof("Starting server at %v", env.GetInt("SERVER_PORT"))
		if err := echoServer.StartServer(httpServer); err != nil {
			log.Logger.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := echoServer.Shutdown(ctx); err != nil {
		return errorutl.Format("error gracefully shutting down server", err)
	}
	return nil
}
