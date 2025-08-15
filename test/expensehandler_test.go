package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fausto4911/expensetracker/internal/handler"
)

func TestGetExpenseByIdHanlder(t *testing.T) {
	app := handler.NewExpenseHandler()

	idExistTest(app, t)

	invalidIdtest(app, t)

	idNotFoundTest(app, t)

}

func idExistTest(app *handler.ExpenseHandler, t *testing.T) {
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

func invalidIdtest(app *handler.ExpenseHandler, t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/expenses/non_int_id", nil)

	rr := httptest.NewRecorder()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/expenses/{id}", app.GetExpenseByIdHanlder)

	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func idNotFoundTest(app *handler.ExpenseHandler, t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/expenses/0000", nil)

	rr := httptest.NewRecorder()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/expenses/{id}", app.GetExpenseByIdHanlder)

	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
