package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/johinsDev/authentication/app/providers"
	"github.com/johinsDev/authentication/handler"
)

func main() {
	// Wrapper logguer
	// wrapper server
	// wrapper session
	// database support multiple driver and support models package to auto load

	store := cookie.NewStore([]byte("secret"))

	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})

	providers.Container().Provide(providers.NewConfig)
	providers.Container().Provide(providers.NewHash)
	providers.Container().Provide(providers.NewDatabase)
	providers.Container().Provide(providers.NewAuth)

	// providers.NewDatabase()

	// providers.NewHash()

	log.Info("Starting server...")

	router := gin.Default()

	router.Use(sessions.Sessions("server-session", store))

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
