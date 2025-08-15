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
	GetAllExpenses() ([]dto.Expense, error)
	CreateExpense(expense dto.Expense) (dto.Expense, error)
	UpdateExpense(expense dto.Expense) (dto.Expense, error)
	DeleteExpenseById(id uint16) error
}

type expenseRepository struct {
	dbConfig config.ExpenseTrackerDBConfig
}

func (r *expenseRepository) CreateExpense(expense dto.Expense) (dto.Expense, error) {
	connUrl := getConnectionUrl(r.dbConfig)
	var err error
	var conn *pgx.Conn
	conn, err = pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return dto.Expense{}, err
	}
	defer conn.Close(context.Background())
	query := `insert into expense (amount, category_id, created) 
	values (@amount, @categoryId, @created)
	RETURNING id`

	args := pgx.NamedArgs{
		"amount":     expense.Amount,
		"categoryId": expense.Category_id,
		"created":    expense.Created,
	}

	var insertedID int

	err = conn.QueryRow(context.Background(), query, args).Scan(&insertedID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running insert: %v\n", err)
		return dto.Expense{}, err
	}
	expense.Id = uint16(insertedID)
	return expense, nil

}

func (r *expenseRepository) GetAllExpenses() ([]dto.Expense, error) {
	connUrl := getConnectionUrl(r.dbConfig)
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return make([]dto.Expense, 0), err
	}
	defer conn.Close(context.Background())
	query := "select * from expense"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return make([]dto.Expense, 0), err
	}
	expenses, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dto.Expense])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return make([]dto.Expense, 0), err
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
	expense.Id = id
	expense.Amount = amount
	expense.Category_id = categoryId
	expense.Created = created
	expense.Modified = modified
	return expense, nil
}

func (r *expenseRepository) DeleteExpenseById(id uint16) error {
	connUrl := getConnectionUrl(r.dbConfig)
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	defer conn.Close(context.Background())
	query := "delete from expense where id = $1"

	res, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (r *expenseRepository) UpdateExpense(expense dto.Expense) (dto.Expense, error) {
	connUrl := getConnectionUrl(r.dbConfig)
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return dto.Expense{}, err
	}
	defer conn.Close(context.Background())
	query := `update expense set amount = @amount, category_id = @categoryId, modified = now()::timestamp where id = @id`
	args := pgx.NamedArgs{
		"amount":     expense.Amount,
		"categoryId": expense.Category_id,
		"created":    expense.Created,
		"id":         expense.Id,
	}

	_, err2 := conn.Exec(context.Background(), query, args)
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "Error running update: %v\n", err2)
		return dto.Expense{}, err2
	}
	return expense, nil
}

func NewExpenseRepository(dbConfig config.ExpenseTrackerDBConfig) ExpenseRepository {
	return &expenseRepository{dbConfig: dbConfig}
}

func getConnectionUrl(conf config.ExpenseTrackerDBConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
}
