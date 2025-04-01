package validator

import (
	"errors"
	"regexp"
)

// ValidateEmail checks if the provided email is valid.
func ValidateEmail(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, email); !matched {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword checks if the provided password meets the criteria.
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

// ValidateExpenseAmount checks if the expense amount is valid.
func ValidateExpenseAmount(amount float64) error {
	if amount <= 0 {
		return errors.New("expense amount must be greater than zero")
	}
	return nil
}