package entity

import (
    "time"
)

type Expense struct {
    ID          int       `json:"id" gorm:"primaryKey"`
    UserID      int       `json:"user_id" gorm:"not null"`
    Amount      float64   `json:"amount" gorm:"not null"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}