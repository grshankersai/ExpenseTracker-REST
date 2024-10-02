package controller

import (
	"github.com/gin-gonic/gin"
	"shanker.com/expense-tracker/middleware"
	"shanker.com/expense-tracker/service"
)

func RegisterRoutes(server *gin.Engine) {
	// Using the logging middleware.
	server.Use(middleware.Logger())

	expenseGroup := server.Group("/expense")
	{
		expenseGroup.GET("", service.GetAllExpenses)       
		expenseGroup.POST("", service.CreateExpense)       
		expenseGroup.PUT("/:id", service.EditExpense)      
		expenseGroup.DELETE("/:id", service.DeleteExpense) 
	}
}