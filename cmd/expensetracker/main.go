package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/Fausto4911/expensetracker/internal/repository"
	"github.com/Fausto4911/expensetracker/internal/service"
)

func main() {

	fmt.Println("expensetracker app started")

	http.HandleFunc("/expenses", expenseHandler)
	http.HandleFunc("GET /expenses/{id}", getExpenseByIdHanlder) //Go 1.22+ Routing Enhancements approach

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func expenseHandler(w http.ResponseWriter, r *http.Request) {

}

func getExpenseByIdHanlder(w http.ResponseWriter, r *http.Request) {
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
