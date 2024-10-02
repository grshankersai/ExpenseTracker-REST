package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "shanker.com/expense-tracker/grpc/expense_pb"
)

func updateExpense(client pb.ExpenseServiceClient, expenseID int64, category string, amount float64, date string, description string) {
	req := &pb.EditExpenseRequest{
		ExpenseId: expenseID,
		Category:  category,
		Amount:    amount,
		Date:      date,
		Description: description,
	}

	resp, err := client.EditExpense(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while updating expense: %v", err)
	}

	fmt.Printf("Updated Expense: %+v\n", resp.Expense)
	fmt.Println("Message from server:", resp.Message)
}

func deleteExpense(client pb.ExpenseServiceClient, expenseID int64) {
	req := &pb.DeleteExpenseRequest{
		ExpenseId: expenseID,
	}

	resp, err := client.DeleteExpense(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while deleting expense: %v", err)
	}

	fmt.Println("Message from server:", resp.Message)
}

func createExpense(client pb.ExpenseServiceClient, category string, amount float64, date string, description string) {
	req := &pb.ExpenseRequest{
		Category:    category,
		Amount:      amount,
		Date:        date,
		Description: description,
	}

	resp, err := client.CreateExpense(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while creating expense: %v", err)
	}

	fmt.Printf("Created Expense: %+v\n", resp.Expense)
	fmt.Println("Message from server:", resp.Message)
}

func getAllExpenses(client pb.ExpenseServiceClient) {
	req := &pb.EmptyRequest{}

	resp, err := client.GetAllExpenses(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while getting all expenses: %v", err)
	}

	fmt.Println("Expenses List:")
	for _, expense := range resp.Expenses {
		fmt.Printf("Expense ID: %d, Category: %s, Amount: %.2f, Date: %s, Description: %s\n",
			expense.ExpenseId, expense.Category, expense.Amount, expense.Date, expense.Description)
	}
}




func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewExpenseServiceClient(conn)	


	
	// createExpense(client, "Groceries", 50.00, "2024-10-02 12:30:00", "Weekly groceries")

	
	getAllExpenses(client)

	// updateExpense(client, 3, "GroceriesEdit", 60, "2024-10-02 12:30:20", "Monthly groceries")

	
	// deleteExpense(client, 10) 

}
