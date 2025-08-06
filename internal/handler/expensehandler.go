package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/Fausto4911/expensetracker/internal/repository"
	"github.com/Fausto4911/expensetracker/internal/service"
)

type ExpenseHandler struct {
}

func (eh *ExpenseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	var res string
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	repo := repository.NewExpenseRepository(dbConfig)
	service := service.NewUserService(repo, dbConfig)

	switch m {
	case "GET":
		fmt.Println(r.URL.Path)
		expense := service.GetExpense(2)
		expenseResponse := dto.BuildExpenseResponse(expense)
		jsonData, err := json.Marshal(expenseResponse)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		data, err := w.Write(jsonData)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		fmt.Println("Write response :: ", data)
	case "POST":
		res = "this is a s default response from POST method."
	case "PUT":
		res = "this a default response from PUT method."
	case "DELETE":
		res = "this a default response from DELETE method."
	default:
		res = "this a default response from NO DETECTED method."
	}
	io.WriteString(w, res)
}

// func getIdFromURL(r *http.Request) string, nil {
// 	// Extract ID from URL path (e.g., /books/1)
// 	idStr := r.URL.Path[len("/books/"):]
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		// http.Error(w, "Invalid book ID", http.StatusBadRequest)
// 		return "", nil
// 	}
// }
