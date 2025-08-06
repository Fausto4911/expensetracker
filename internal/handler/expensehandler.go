package handler

import (
	"io"
	"net/http"

	"github.com/Fausto4911/expensetracker/internal/service"
)

type ExpenseHandler struct {
	es service.ExpenseService
}

func (eh *ExpenseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	var res string
	eh.es = service.ExpenseService{}
	switch m {
	case "GET":
		eh.es.GetExpenseById(2)
		res = "thi is a default response from GET method."
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
