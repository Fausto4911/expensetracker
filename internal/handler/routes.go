package handler

import "net/http"

func (eh *ExpenseHandler) Routes() *http.ServeMux {

	// Declare a new servemux and add a /v1/healthcheck route which dispatches requests
	// to the healthcheckHandler method (which we will create in a moment).
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthcheck", eh.HealthcheckHandler)
	mux.HandleFunc("GET /v1/expenses", eh.GetAllExpenses)
	mux.HandleFunc("GET /v1/expenses/{id}", eh.GetExpenseByIdHanlder) //Go 1.22+ Routing Enhancements approach
	mux.HandleFunc("POST /v1/expenses", eh.CreateExpenseHandler)
	mux.HandleFunc("DELETE /v1/expenses/{id}", eh.DeleteExpense)
	mux.HandleFunc("PUT /v1/expenses/{id}", eh.UpdateExpense)

	return mux
}
