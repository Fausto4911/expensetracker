package service

import (
	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/Fausto4911/expensetracker/internal/repository"
)

type ExpenseService interface {
	GetExpense(id uint16) (dto.Expense, error)
	GetAllExpenses() ([]dto.Expense, error)
	CreateExpense(expense dto.Expense) (dto.Expense, error)
	UpdateExpense(expense dto.Expense) (dto.Expense, error)
	DeleteExpenseById(id uint16) error
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

func (es expenseService) GetAllExpenses() ([]dto.Expense, error) {
	expenses, err := es.repo.GetAllExpenses()
	if err != nil {
		return make([]dto.Expense, 0), err
	}
	return expenses, nil

}

func (es expenseService) CreateExpense(expense dto.Expense) (dto.Expense, error) {
	expense, err := es.repo.CreateExpense(expense)
	if err != nil {
		return dto.Expense{}, err
	}
	return expense, nil
}

func (es expenseService) DeleteExpenseById(id uint16) error {
	err := es.repo.DeleteExpenseById(id)
	if err != nil {
		return err
	}
	return nil

}

func (es expenseService) UpdateExpense(expense dto.Expense) (dto.Expense, error) {
	expense, err := es.repo.UpdateExpense(expense)
	if err != nil {
		return dto.Expense{}, err
	}
	return expense, nil
}
