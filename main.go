package main

import (
	"net/http"
	"os"
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

	log.Info().Str("addr", server.Addr).Msg("starting server")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.ListStatus)
	router.HandleFunc("/{code}", handler.GetStatus)

	return router
}
