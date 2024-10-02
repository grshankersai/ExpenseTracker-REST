package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"shanker.com/expense-tracker/db"
	"shanker.com/expense-tracker/entity"
	pb "shanker.com/expense-tracker/grpc/expense_pb"
	"shanker.com/expense-tracker/repository"
)

type ExpenseServer struct {
	pb.UnimplementedExpenseServiceServer
}

func (s *ExpenseServer) GetAllExpenses(ctx context.Context, req *pb.EmptyRequest) (*pb.ExpenseListResponse, error) {

	
	
	expenses, err := repository.GetAllExpenses()

	
	if err != nil {
		return nil, fmt.Errorf("error fetching expenses: %v", err)
	}
	
	var expenseList []*pb.Expense
	for _, e := range expenses {
		expenseList = append(expenseList, &pb.Expense{
			ExpenseId:   e.ExpenseID,
			Category:    e.Category,
			Amount:      float64(e.Amount),
			Date:        e.Date.String(),
			Description: e.Description,
		})
	}

	return &pb.ExpenseListResponse{Expenses: expenseList}, nil
}

func (s *ExpenseServer) CreateExpense(ctx context.Context, req *pb.ExpenseRequest) (*pb.ExpenseResponse, error) {

	
	date, err := time.Parse("2006-01-02 15:04:05", req.GetDate())
	if err != nil {
		return nil, err 
	}

	expense := entity.Expense{
		Category:    req.GetCategory(),
		Amount:      float64(req.GetAmount()),
		Date:        date,
		Description: req.GetDescription(),
	}
	err = repository.Save(&expense)	

	if err != nil {
		return nil, fmt.Errorf("failed to create expense: %v", err)
	}

	return &pb.ExpenseResponse{Expense: &pb.Expense{
		ExpenseId:   expense.ExpenseID,
		Category:    expense.Category,
		Amount:      float64(expense.Amount),
		Date:        expense.Date.Format("2006-01-02 15:04:05"),
		Description: expense.Description,
	}, Message: "Expense created!"}, nil
}

func (s *ExpenseServer) EditExpense(ctx context.Context, req *pb.EditExpenseRequest) (*pb.ExpenseResponse, error) {
	
	date, err := time.Parse("2006-01-02 15:04:05", req.GetDate())
	if err != nil {
		return nil, err 
	}

	expense := entity.Expense{
		ExpenseID:   req.GetExpenseId(),
		Category:    req.GetCategory(),
		Amount:      float64(req.GetAmount()),
		Date:        date,
		Description: req.GetDescription(),
	}

	err = repository.Modify(req.GetExpenseId(), &expense)
	if err != nil {
		return nil, fmt.Errorf("failed to edit expense: %v", err)
	}

	return &pb.ExpenseResponse{Expense: &pb.Expense{
		ExpenseId:   expense.ExpenseID,
		Category:    expense.Category,
		Amount:      float64(expense.Amount),
		Date:        expense.Date.Format("2006-01-02 15:04:05"),
		Description: expense.Description,
	}, Message: "Expense updated!"}, nil
}

func (s *ExpenseServer) DeleteExpense(ctx context.Context, req *pb.DeleteExpenseRequest) (*pb.DeleteResponse, error) {
	err := repository.Delete(req.GetExpenseId())
	if err != nil {
		return nil, fmt.Errorf("failed to delete expense: %v", err)
	}

	return &pb.DeleteResponse{Message: "Expense deleted!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	db.InitDB()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExpenseServiceServer(s, &ExpenseServer{})

	log.Printf("gRPC server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
