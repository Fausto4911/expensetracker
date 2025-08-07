package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fausto4911/expensetracker/internal/handler"
)

func main() {

	fmt.Println("expensetracker app started")
	eh := &handler.ExpenseHandler{}

	http.HandleFunc("GET /expenses", eh.GetAllExpenses)
	http.HandleFunc("GET /expenses/{id}", eh.GetExpenseByIdHanlder) //Go 1.22+ Routing Enhancements approach

	log.Fatal(http.ListenAndServe(":8080", nil))

}
