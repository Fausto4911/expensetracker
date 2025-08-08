package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/Fausto4911/expensetracker/internal/repository"
	"github.com/Fausto4911/expensetracker/internal/service"
)

const version = "1.0.0"

// Define a config struct to hold all the configuration settings for our application.
type Config struct {
	Port int
	Env  string
}

// Define an application struct to hold the dependencies for our HTTP handlers, helpers,and middleware.
type ExpenseHandler struct {
	Config Config
	Logger *slog.Logger
}

func (eh *ExpenseHandler) GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	eh.Logger.Info("GetAllExpenses - started")
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	repo := repository.NewExpenseRepository(dbConfig)
	service := service.NewExpenseService(repo, dbConfig)
	expenses, err := service.GetAllExpenses()
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error parsing ID ", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(expenses)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error marshalling JSON:", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := w.Write(jsonData)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error writing data:", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	eh.Logger.Info("GetAllExpenses - ", slog.Int("data", data))
	eh.Logger.Info("GetAllExpenses - finish")
}

func (eh *ExpenseHandler) GetExpenseByIdHanlder(w http.ResponseWriter, r *http.Request) {
	eh.Logger.Info("GetExpenseByIdHanlder started")
	id := r.PathValue("id")
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	repo := repository.NewExpenseRepository(dbConfig)
	service := service.NewExpenseService(repo, dbConfig)

	u, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error parsing ID ", http.StatusInternalServerError)
		return
	}
	expense, err := service.GetExpense(uint16(u))
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error getting expense ", http.StatusInternalServerError)
		return
	}
	expenseResponse := dto.BuildExpenseResponse(expense)
	jsonData, err := json.Marshal(expenseResponse)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error marshalling JSON:", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := w.Write(jsonData)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error writing data:", http.StatusInternalServerError)
		return
	}
	eh.Logger.Info("GetAllExpenses - ", slog.Int("data", data))
	eh.Logger.Info("GetExpenseByIdHanlder finish")
}

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (eh *ExpenseHandler) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	eh.Logger.Info("HealthcheckHandler started")
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", eh.Config.Env)
	fmt.Fprintf(w, "version: %s\n", version)
	eh.Logger.Info("HealthcheckHandler finsh")
}
