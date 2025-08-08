package service

import (
	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/Fausto4911/expensetracker/internal/repository"
)

type ExpenseService interface {
	GetExpense(id uint16) (dto.Expense, error)
	GetAllExpenses() ([]dto.ExpenseResponse, error)
	CreateExpense(expense dto.ExpenseResponse) (dto.ExpenseResponse, error)
	// UpdateExpense(expense dto.Expense) dto.Expense
	// DeleteExpenseById(id uint16) bool
}

type expenseService struct {
	repo     repository.ExpenseRepository
	dbConfig config.ExpenseTrackerDBConfig
}

func NewExpenseService(repo repository.ExpenseRepository, dbConfig config.ExpenseTrackerDBConfig) ExpenseService {
	return &expenseService{repo: repo, dbConfig: dbConfig}
}

func (es expenseService) GetExpense(id uint16) (dto.Expense, error) {
	expense, err := es.repo.GetExpenseById(id)
	if err != nil {
		return dto.Expense{}, err
	}
	return expense, nil
}

func (es expenseService) GetAllExpenses() ([]dto.ExpenseResponse, error) {
	expenses, err := es.repo.GetAllExpenses()
	if err != nil {
		return make([]dto.ExpenseResponse, 0), err
	}
	return expenses, nil

}

func (es expenseService) CreateExpense(expense dto.ExpenseResponse) (dto.ExpenseResponse, error) {
	expense, err := es.repo.CreateExpense(expense)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}
	return expense, nil
}
