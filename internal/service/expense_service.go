package service

import (
	"context"
	"fmt"
	"os"

	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ExpenseService struct {
}

func (es ExpenseService) GetExpenseById(id uint16) {
	connUrl := getConnectionUrl(config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"})
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
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
		os.Exit(1)
	}

	fmt.Printf("id : %d | amount : %f | category_id: %d | created: %s | modified: %s", expenseId, amount, categoryId, created, modified)
}

func getConnectionUrl(conf config.ExpenseTrackerDBConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
}
