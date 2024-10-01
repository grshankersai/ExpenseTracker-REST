package validations

import (
	"errors"

	"shanker.com/expense-tracker/entity"
)

func ValidateExpenseObject(expense *entity.Expense) error{
	
	if expense.Amount <= 0 {		
		return errors.New("amount should be greater than zero")
	}

	
	if expense.Category == "" {		
		return errors.New("amount should be greater than zero")
	}

	return nil

}