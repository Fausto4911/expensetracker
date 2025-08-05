package handler

import (
	"io"
	"net/http"
)

type ExpenseHandler struct {
}

func (e *ExpenseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	var res string
	switch m {
	case "GET":
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
