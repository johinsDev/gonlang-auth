package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gookit/config"
	"github.com/gookit/config/yaml"
	"github.com/johinsDev/authentication/handler"
)

func main() {
	// Init config
	// init dig
	// logger wrap logrus, provier
	// databa providers
	// inject models
	// gin provider
	// readme files ( app/dto, app/models app/controllers app/routes app/middleware app/services app/repository, config, boostrap, mails, lib)

	// you could insert your favorite logger here for structured or leveled logging
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})

	log.Info("Starting server...")

	config.WithOptions(config.ParseEnv)

	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config/auth.yml")

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	handler.NewHandler(&handler.Config{
		R: router,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Failed to initialize server: %v\n", err)
		}
	}()

	log.Info("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Info("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown: %v\n", err)
	}
}
