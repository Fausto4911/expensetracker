package handler

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/Fausto4911/expensetracker/internal/repository"
	"github.com/Fausto4911/expensetracker/internal/service"
	"github.com/jackc/pgx/v5"
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

func NewExpenseHandler() *ExpenseHandler {
	var cfg Config

	// Read the value of the port and env command-line flags into the config struct. We
	// default to using the port number 8080 and the environment "development" if no
	// corresponding flags are provided.
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Declare an instance of the ExpenseHandler struct, containing the config struct and
	// the logger.
	app := &ExpenseHandler{
		Config: cfg,
		Logger: logger,
	}
	return app
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
	eh.Logger.Info(fmt.Sprintf("GetExpenseByIdHanlder id ==> %v", id))
	eh.Logger.Info(r.URL.Path)
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	repo := repository.NewExpenseRepository(dbConfig)
	service := service.NewExpenseService(repo, dbConfig)

	u, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Wrong id format ", http.StatusBadRequest)
		return
	}
	expense, err := service.GetExpense(uint16(u))
	if err != nil {
		if err == sql.ErrNoRows || err == pgx.ErrNoRows {
			errString := fmt.Sprintf("No data found for id : %v", uint16(u))
			eh.Logger.Error(errString)
			http.Error(w, errString, http.StatusNotFound)
			return
		}
		eh.Logger.Error(err.Error())
		http.Error(w, "Error getting expense ", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(expense)
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

func (eh *ExpenseHandler) CreateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	eh.Logger.Info("CreateExpenseHandler started")
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	repo := repository.NewExpenseRepository(dbConfig)
	service := service.NewExpenseService(repo, dbConfig)
	var err error
	var expense dto.Expense
	body, err := io.ReadAll(r.Body)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error Reading body", http.StatusInternalServerError)
		return
	}
	expense = dto.Expense{}
	err = json.Unmarshal(body, &expense)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error while Unmarshal body", http.StatusInternalServerError)
		return
	}

	expense, err = service.CreateExpense(expense)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error while Creating Expense", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Expense :: ==> %s", expense)
	fmt.Printf("%s", body)

}

func (eh *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	eh.Logger.Info("DeleteExpense started")
	id := r.PathValue("id")
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	repo := repository.NewExpenseRepository(dbConfig)
	service := service.NewExpenseService(repo, dbConfig)
	var err error

	u, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error parsing ID ", http.StatusInternalServerError)
		return
	}
	err = service.DeleteExpenseById(uint16(u))
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error while Creating Expense", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (eh *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	eh.Logger.Info("UpdateExpense started")
	id := r.PathValue("id")
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	repo := repository.NewExpenseRepository(dbConfig)
	service := service.NewExpenseService(repo, dbConfig)
	var err error

	u, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error parsing ID ", http.StatusInternalServerError)
		return
	}
	fmt.Println("id ", u)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error Reading body", http.StatusInternalServerError)
		return
	}
	expense := dto.Expense{}
	err = json.Unmarshal(body, &expense)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error while Unmarshal body", http.StatusInternalServerError)
		return
	}
	expense.Id = uint16(u)

	expense, err = service.UpdateExpense(expense)
	if err != nil {
		eh.Logger.Error(err.Error())
		http.Error(w, "Error while Updating Expense", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Expense :: ==> %s", expense)
	fmt.Printf("%s", body)

}
