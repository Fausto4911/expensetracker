package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/Fausto4911/expensetracker/internal/handler"
	"github.com/Fausto4911/expensetracker/internal/repository"
	"github.com/Fausto4911/expensetracker/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestGetExpenseByIdHanlder(t *testing.T) {
	app := handler.NewExpenseHandler()

	category, expense := setDataForTest(t)

	fmt.Println(category.Id)

	idExistTest(app, t, expense.Id)

	invalidIdtest(app, t)

	idNotFoundTest(app, t)

}

func idExistTest(app *handler.ExpenseHandler, t *testing.T, expenseId uint16) {
	fmt.Printf("expense ID ==> %v", expenseId)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/expenses/%v", expenseId), nil)

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

func setDataForTest(t *testing.T) (c dto.Category, e dto.Expense) {
	dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	categoryRepo := repository.NewCategoryRepository(dbConfig)
	categorySrv := service.NewCategoryService(categoryRepo, dbConfig)
	var err error
	var category dto.Category
	category.Name = "TestCategory"
	category.Description = "TestDescription"

	category, err = categorySrv.CreateCategory(category)
	if err != nil {
		t.Fatal("Error creating category test data")
	}

	expenseRepo := repository.NewExpenseRepository(dbConfig)
	expenseSrv := service.NewExpenseService(expenseRepo, dbConfig)
	var expense dto.Expense
	expense.Amount = 120
	expense.Category_id = category.Id
	// Get the current time in the local timezone
	currentTime := time.Now()

	// Create a pgtype.Timestamptz from the current time
	currentTimestamptz := pgtype.Timestamptz{
		Time:  currentTime,
		Valid: true, // Indicate that the value is present (not null)
	}
	expense.Created = currentTimestamptz

	fmt.Println(expense)

	expense, err = expenseSrv.CreateExpense(expense)

	if err != nil {
		fmt.Println(err)
		t.Fatal("Error creating expense test data")

	}

	return category, expense

}
