package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Fausto4911/expensetracker/internal/handler"
)

func main() {

	var cfg handler.Config

	// Read the value of the port and env command-line flags into the config struct. We
	// default to using the port number 8080 and the environment "development" if no
	// corresponding flags are provided.
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Declare an instance of the ExpenseHandler struct, containing the config struct and
	// the logger.
	app := &handler.ExpenseHandler{
		Config: cfg,
		Logger: logger,
	}

	// Declare an HTTP server which listens on the port provided in the config struct,
	// uses the servemux we created above as the handler, has some sensible timeout
	// settings, and writes any log messages to the structured logger at Error level.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Start the HTTP server.
	logger.Info("starting server", "addr", srv.Addr, "env", cfg.Env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}
