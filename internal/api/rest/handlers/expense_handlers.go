package handlers

import (
	"net/http"

	"expense-management-system/internal/db/models"
	"expense-management-system/internal/db/repository"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	repo repository.ExpenseRepository
}

func NewExpenseHandler(repo repository.ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: repo}
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create expense"})
		return
	}

	c.JSON(http.StatusCreated, expense)
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	expenses, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve expenses"})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *ExpenseHandler) GetExpenseByID(c *gin.Context) {
	id := c.Param("id")
	expense, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	id := c.Param("id")
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Update(id, &expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update expense"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete expense"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
