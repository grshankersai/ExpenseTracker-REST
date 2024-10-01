package entity

import (
	"shanker.com/expense-tracker/db"
)

const getAllRecordsQuery = "SELECT * FROM expenses"

const insertRecordQuery =  "INSERT INTO expenses(category, amount, date, description) VALUES ($1, $2, $3, $4) RETURNING expense_id"

const updateRecordQuery = "UPDATE expenses SET category = $1, amount = $2, date = $3, description = $4 	WHERE expense_id = $5"

const deleteRecordQuery = "DELETE FROM expenses	WHERE expense_id = $1"




func GetAllExpenses() ([]Expense, error) {

	stmt, err := db.DB.Prepare(getAllRecordsQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
		
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense

	for rows.Next() {
		var expense Expense
		err := rows.Scan(&expense.ExpenseID, &expense.Category, &expense.Amount, &expense.Date, &expense.Description)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}



func (expense *Expense) Save() error {

	stmt, err := db.DB.Prepare(insertRecordQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(expense.Category, expense.Amount, expense.Date, expense.Description).Scan(&expense.ExpenseID)
	if err != nil {
		return err
	}

	return nil
}


func (expense *Expense) Modify(eid int64) error {

	stmt, err := db.DB.Prepare(updateRecordQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(expense.Category, expense.Amount, expense.Date, expense.Description, eid)
	if err != nil {
		return err
	}

	return nil
}


func Delete(eid int64) error {

	stmt, err := db.DB.Prepare(deleteRecordQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(eid)
	if err != nil {
		return err
	}

	return nil
}

