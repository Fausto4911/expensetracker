package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Fausto4911/expensetracker/internal/handler"
)

func main() {
	app := handler.NewExpenseHandler()

	// Declare an HTTP server which listens on the port provided in the config struct,
	// uses the servemux we created above as the handler, has some sensible timeout
	// settings, and writes any log messages to the structured logger at Error level.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelError),
	}

	// Start the HTTP server.
	app.Logger.Info("starting server", "addr", srv.Addr, "env", app.Config.Env)

	err := srv.ListenAndServe()
	app.Logger.Error(err.Error())
	os.Exit(1)

}
