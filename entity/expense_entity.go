package entity

import (
	"time"
)

type Expense struct {
	ExpenseID   int64      `json:"expense_id" db:"expense_id"` 
	Category    string    `json:"category" db:"category"`         
	Amount      float64   `json:"amount" db:"amount"`             
	Date        time.Time `json:"date" db:"date"`                 
	Description string    `json:"description,omitempty" db:"description"` 
}