package service

import (
	"errors"
	"expense-management-system/internal/db/models"
	"expense-management-system/internal/db/repository"
)

type ExpenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) CreateExpense(expense *models.Expense) error {
	if expense.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return s.repo.Create(expense)
}

func (s *ExpenseService) GetExpenseByID(id int) (*models.Expense, error) {
	return s.repo.FindByID(id)
}

func (s *ExpenseService) UpdateExpense(expense *models.Expense) error {
	if expense.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return s.repo.Update(expense)
}

func (s *ExpenseService) DeleteExpense(id int) error {
	return s.repo.Delete(id)
}

func (s *ExpenseService) ListExpenses() ([]models.Expense, error) {
	return s.repo.FindAll()
}
