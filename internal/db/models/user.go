package models

import (
    "time"
)

type User struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Username  string `gorm:"unique;not null" json:"username"`
    Password  string `gorm:"not null" json:"-"`
    Email     string `gorm:"unique;not null" json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}