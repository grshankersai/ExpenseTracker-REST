package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=expensedb sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createExpenseTable := `
	CREATE TABLE IF NOT EXISTS expenses (
	expense_id SERIAL PRIMARY KEY,
	category TEXT NOT NULL,
	amount NUMERIC(10, 2) NOT NULL,
	date TIMESTAMP NOT NULL,
	description TEXT
)`

	_, err := DB.Exec(createExpenseTable)
	if err != nil {
		panic("Couldn't create expense table: " + err.Error())
	}

	fmt.Println("Expense table created successfully")
}
