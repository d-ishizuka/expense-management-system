package entity

import (
    "time"
)

type Payment struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Amount      float64   `json:"amount"`
    Method      string    `json:"method"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}