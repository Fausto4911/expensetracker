package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ExpenseRepository interface {
	GetExpenseById(id uint16) (dto.Expense, error)
	GetAllExpenses() ([]dto.ExpenseResponse, error)
	// CreateExpense(expense dto.Expense) dto.Expense
	// UpdateExpense(expense dto.Expense) dto.Expense
	// DeleteExpenseById(id uint16) bool
}

type expenseRepository struct {
	dbConfig config.ExpenseTrackerDBConfig
}

func (r *expenseRepository) GetAllExpenses() ([]dto.ExpenseResponse, error) {
	connUrl := getConnectionUrl(r.dbConfig)
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return make([]dto.ExpenseResponse, 0), err
	}
	defer conn.Close(context.Background())
	query := "select * from expense"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return make([]dto.ExpenseResponse, 0), err
	}
	expenses, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dto.ExpenseResponse])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return make([]dto.ExpenseResponse, 0), err
	}
	fmt.Println(expenses)
	return expenses, nil
}

func (r *expenseRepository) GetExpenseById(id uint16) (dto.Expense, error) {
	connUrl := getConnectionUrl(r.dbConfig)
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return dto.Expense{}, err
	}
	defer conn.Close(context.Background())

	var (
		expenseId  uint16
		amount     float32
		categoryId uint16
		created    pgtype.Timestamptz
		modified   pgtype.Timestamptz
	)

	query := "select * from expense where id = $1"
	err = conn.QueryRow(context.Background(), query, id).Scan(&expenseId, &amount, &categoryId, &created, &modified)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return dto.Expense{}, err
		// os.Exit(1)
	}

	fmt.Printf("id : %d | amount : %f | category_id: %d | created: %s | modified: %s\n", expenseId, amount, categoryId, created, modified)
	expense := dto.Expense{}
	expense.SetId(id)
	expense.SetAmount(amount)
	expense.SetCategoryId(categoryId)
	expense.SetCreated(created)
	expense.SetModified(modified)
	return expense, nil
}

func NewExpenseRepository(dbConfig config.ExpenseTrackerDBConfig) ExpenseRepository {
	return &expenseRepository{dbConfig: dbConfig}
}

func getConnectionUrl(conf config.ExpenseTrackerDBConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
}
