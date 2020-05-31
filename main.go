package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/maximepeschard/statuscode/handler"
	"github.com/maximepeschard/statuscode/middleware"
	"github.com/maximepeschard/statuscode/status"
	"github.com/rs/zerolog/log"
)

func main() {
	var err error

	port := os.Getenv("PORT")
	if port == "" {
		log.Warn().Msg("no $PORT, using default port")
		port = "8080"
	}

	err = status.Init("data.toml")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize")
	}

	router := routes()
	router.Use(middleware.Logging)
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Server runs in a goroutine so that it does not block.
	go func() {
	log.Info().Str("addr", server.Addr).Msg("starting server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Send()
		}
	}()

	// We handle graceful shutdowns for SIGINT and SIGTERM.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info().Msg("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.ListStatus)
	router.HandleFunc("/{code}", handler.GetStatus)

	return router
}
