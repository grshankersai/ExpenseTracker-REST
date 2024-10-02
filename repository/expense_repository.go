package repository

import (
	"errors"

	"shanker.com/expense-tracker/db"
	"shanker.com/expense-tracker/entity"
)

const getAllRecordsQuery = "SELECT * FROM expenses"

const insertRecordQuery =  "INSERT INTO expenses(category, amount, date, description) VALUES ($1, $2, $3, $4) RETURNING expense_id"

const updateRecordQuery = "UPDATE expenses SET category = $1, amount = $2, date = $3, description = $4 	WHERE expense_id = $5"

const deleteRecordQuery = "DELETE FROM expenses	WHERE expense_id = $1"

const checkIfRecordExists = "SELECT EXISTS(SELECT 1 FROM expenses WHERE expense_id = $1)"




func GetAllExpenses() ([]entity.Expense, error) {

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

	var expenses []entity.Expense

	for rows.Next() {
		var expense entity.Expense
		err := rows.Scan(&expense.ExpenseID, &expense.Category, &expense.Amount, &expense.Date, &expense.Description)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}



func Save(expense *entity.Expense) error {

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


func  Modify(eid int64,expense *entity.Expense) error {
	err := checkIfExists(eid)

	if(err != nil){
		return err
	}
	
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

func checkIfExists(eid int64) error{
	stmtCheck, err := db.DB.Prepare(checkIfRecordExists)
	if err != nil {
		return err 
	}
	defer stmtCheck.Close()

	
	var exists bool
	err = stmtCheck.QueryRow(eid).Scan(&exists)
	if err != nil {
		return err 
	}

	if(!exists){
		return errors.New("the expense doesnot exists")
	}
	return nil;
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

