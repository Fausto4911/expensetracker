package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/Fausto4911/expensetracker/internal/repository"
	"github.com/Fausto4911/expensetracker/internal/service"
)

type ExpenseHandler struct {
}

func (eh *ExpenseHandler) GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllExpenses started")
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	repo := repository.NewExpenseRepository(dbConfig)
	service := service.NewExpenseService(repo, dbConfig)
	expenses, err := service.GetAllExpenses()
	if err != nil {
		http.Error(w, "Error parsing ID ", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(expenses)
	if err != nil {
		http.Error(w, "Error marshalling JSON:", http.StatusInternalServerError)
		return
	}
	data, err := w.Write(jsonData)
	if err != nil {
		http.Error(w, "Error writing data:", http.StatusInternalServerError)
		return
	}
	fmt.Println("Write response :: ", data)
}

func (eh *ExpenseHandler) GetExpenseByIdHanlder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	repo := repository.NewExpenseRepository(dbConfig)
	service := service.NewExpenseService(repo, dbConfig)

	u, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		http.Error(w, "Error parsing ID ", http.StatusInternalServerError)
		return
	}
	expense, err := service.GetExpense(uint16(u))
	if err != nil {
		http.Error(w, "Error getting expense ", http.StatusInternalServerError)
		return
	}
	expenseResponse := dto.BuildExpenseResponse(expense)
	jsonData, err := json.Marshal(expenseResponse)
	if err != nil {
		http.Error(w, "Error marshalling JSON:", http.StatusInternalServerError)
		return
	}
	data, err := w.Write(jsonData)
	if err != nil {
		http.Error(w, "Error writing data:", http.StatusInternalServerError)
		return
	}
	fmt.Println("Write response :: ", data)
}
