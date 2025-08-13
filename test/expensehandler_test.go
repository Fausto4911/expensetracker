package test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Fausto4911/expensetracker/internal/handler"
)

func TestGetExpenseByIdHanlder(t *testing.T) {
	app := &handler.ExpenseHandler{
		Config: handler.Config{
			Port: 8080,
			Env:  "development",
		},
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/expenses/4", nil)

	rr := httptest.NewRecorder()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/expenses/{id}", app.GetExpenseByIdHanlder)

	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
