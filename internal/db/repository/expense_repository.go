package repository

import (
	"expense-management-system/internal/db/models"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Create(expense *models.Expense) error
	GetByID(id uint) (*models.Expense, error)
	GetAll() ([]models.Expense, error)
	Update(expense *models.Expense) error
	Delete(id uint) error
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) Create(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepository) GetByID(id uint) (*models.Expense, error) {
	var expense models.Expense
	if err := r.db.First(&expense, id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *expenseRepository) GetAll() ([]models.Expense, error) {
	var expenses []models.Expense
	if err := r.db.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (r *expenseRepository) Update(expense *models.Expense) error {
	return r.db.Save(expense).Error
}

func (r *expenseRepository) Delete(id uint) error {
	return r.db.Delete(&models.Expense{}, id).Error
}
