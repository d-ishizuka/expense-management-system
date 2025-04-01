package resolvers

import (
	"context"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/service"
)

type ExpenseResolver struct {
	expenseService service.ExpenseService
}

func NewExpenseResolver(expenseService service.ExpenseService) *ExpenseResolver {
	return &ExpenseResolver{expenseService: expenseService}
}

func (r *ExpenseResolver) GetExpenses(ctx context.Context) ([]entity.Expense, error) {
	return r.expenseService.GetAllExpenses(ctx)
}

func (r *ExpenseResolver) CreateExpense(ctx context.Context, input entity.Expense) (entity.Expense, error) {
	return r.expenseService.CreateExpense(ctx, input)
}

func (r *ExpenseResolver) UpdateExpense(ctx context.Context, id string, input entity.Expense) (entity.Expense, error) {
	return r.expenseService.UpdateExpense(ctx, id, input)
}

func (r *ExpenseResolver) DeleteExpense(ctx context.Context, id string) (bool, error) {
	return r.expenseService.DeleteExpense(ctx, id)
}
