package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/jackc/pgx/v5"
)

type CategoryRepository interface {
	CreateCategory(expense dto.Category) (dto.Category, error)
}

type categoryRepository struct {
	dbConfig config.ExpenseTrackerDBConfig
}

func (r *categoryRepository) CreateCategory(category dto.Category) (dto.Category, error) {
	connUrl := getConnectionUrl(r.dbConfig)
	var err error
	var conn *pgx.Conn
	conn, err = pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return dto.Category{}, err
	}
	defer conn.Close(context.Background())
	query := `insert into category (name, description) 
	values (@name, @description)
	RETURNING id`
	args := pgx.NamedArgs{
		"name":        category.Name,
		"description": category.Description,
	}

	var insertedID int

	err = conn.QueryRow(context.Background(), query, args).Scan(&insertedID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running insert: %v\n", err)
		return dto.Category{}, err
	}
	category.Id = uint16(insertedID)
	return category, nil

}

func NewCategoryRepository(dbConfig config.ExpenseTrackerDBConfig) CategoryRepository {
	return &categoryRepository{dbConfig: dbConfig}
}
