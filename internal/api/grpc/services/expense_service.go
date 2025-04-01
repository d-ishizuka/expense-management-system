package services

import (
	"context"
	"expense-management-system/internal/db/models"
	"expense-management-system/internal/db/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExpenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) CreateExpense(ctx context.Context, expense *models.Expense) (*models.Expense, error) {
	if err := s.repo.Create(expense); err != nil {
		return nil, status.Errorf(codes.Internal, "could not create expense: %v", err)
	}
	return expense, nil
}

func (s *ExpenseService) GetExpense(ctx context.Context, id string) (*models.Expense, error) {
	expense, err := s.repo.GetByID(id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "expense not found: %v", err)
	}
	return expense, nil
}

func (s *ExpenseService) UpdateExpense(ctx context.Context, expense *models.Expense) (*models.Expense, error) {
	if err := s.repo.Update(expense); err != nil {
		return nil, status.Errorf(codes.Internal, "could not update expense: %v", err)
	}
	return expense, nil
}

func (s *ExpenseService) DeleteExpense(ctx context.Context, id string) error {
	if err := s.repo.Delete(id); err != nil {
		return status.Errorf(codes.Internal, "could not delete expense: %v", err)
	}
	return nil
}
