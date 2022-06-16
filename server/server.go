package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-todo/db/mysql"
	"go-todo/server/api"
	"go-todo/server/config"
	errorutl "go-todo/utl/error"
	"go-todo/utl/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start(cfg *config.Configuration) {
	db, err := mysql.Connect(cfg)
	errorutl.Fatal(err)

	echoServer := api.Start(cfg, db)

	startServer(cfg, echoServer)
}

func startServer(cfg *config.Configuration, echoServer *echo.Echo) {

	httpServer := &http.Server{
		Addr:         cfg.Server.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeoutSeconds) * time.Second,
	}

	echoServer.Debug = cfg.Server.Debug

	// Start server
	go func() {
		log.Logger.Info("Starting server...")
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
		errorutl.Fatal(err)
	}
}
