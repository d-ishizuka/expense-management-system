package db

import (
	"expense-management-system/internal/db/models"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	// ユーザーのシードデータ
	users := []models.User{
		{Name: "山田太郎", Email: "taro@example.com"},
		{Name: "鈴木花子", Email: "hanako@example.com"},
	}

	// 経費のシードデータ
	expenses := []models.Expense{
		{UserID: 1, Amount: 1000, Description: "ランチ"},
		{UserID: 1, Amount: 2000, Description: "交通費"},
		{UserID: 2, Amount: 1500, Description: "会議費"},
	}

	// 支払いのシードデータ
	payments := []models.Payment{
		{UserID: 1, Amount: 1000, Method: "クレジットカード"},
		{UserID: 2, Amount: 1500, Method: "現金"},
	}

	// データベースにシードデータを挿入
	if err := db.Create(&users).Error; err != nil {
		return err
	}
	if err := db.Create(&expenses).Error; err != nil {
		return err
	}
	if err := db.Create(&payments).Error; err != nil {
		return err
	}

	return nil
}
