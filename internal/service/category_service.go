package service

import (
	"github.com/Fausto4911/expensetracker/internal/config"
	"github.com/Fausto4911/expensetracker/internal/dto"
	"github.com/Fausto4911/expensetracker/internal/repository"
)

type CategoryService interface {
	CreateCategory(dto.Category) (dto.Category, error)
}

type categoryService struct {
	repo     repository.CategoryRepository
	dbConfig config.ExpenseTrackerDBConfig
}

func NewCategoryService(repo repository.CategoryRepository, dbConfig config.ExpenseTrackerDBConfig) CategoryService {
	return &categoryService{repo: repo, dbConfig: dbConfig}
}

func (service *categoryService) CreateCategory(c dto.Category) (dto.Category, error) {
	category, err := service.repo.CreateCategory(c)
	if err != nil {
		return dto.Category{}, err
	}
	return category, nil
}
