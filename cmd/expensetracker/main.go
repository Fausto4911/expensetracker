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
	http.Handle("/expenses", eh)
	http.Handle("/expenses/{id}", eh)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
