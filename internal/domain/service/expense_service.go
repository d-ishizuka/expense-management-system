package service

import (
	"errors"
	"expense-management-system/internal/db/models"
	"expense-management-system/internal/db/repository"
	"expense-management-system/internal/domain/entity"
)

type ExpenseService struct {
	repo repository.ExpenseRepository
}

func (s *ExpenseService) GetAllExpenses() ([]*entity.Expense, error) {
	expenseModel, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// モデルからエンティティへの変換
	expenses := make([]*entity.Expense, 0, len(expenseModel))
	for _, model := range expenseModel {
		expense := &entity.Expense{
			ID: model.ID,
			UserID: model.UserID,
			Amount: model.Amount,
			Description: model.Description,
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
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

func (s *ExpenseService) GetExpenseByID(id uint) (*models.Expense, error) {
	return s.repo.GetByID(id)
}

func (s *ExpenseService) UpdateExpense(expense *models.Expense) error {
	if expense.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return s.repo.Update(expense)
}

func (s *ExpenseService) DeleteExpense(id uint) error {
	return s.repo.Delete(id)
}

func (s *ExpenseService) ListExpenses() ([]models.Expense, error) {
	return s.repo.GetAll()
}
