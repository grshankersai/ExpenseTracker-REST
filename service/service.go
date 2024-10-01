package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"shanker.com/expense-tracker/entity"
	"shanker.com/expense-tracker/validations"
)

func GetAllExpenses(context *gin.Context) {
	expenses , err := entity.GetAllExpenses()
	
	if err!=nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldnt Fetch!!"})
		return
	}

	context.JSON(http.StatusOK,expenses)
}



func CreateExpense(context *gin.Context) {
	var expense entity.Expense
	err := context.ShouldBindJSON(&expense)

	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request"})
		return
	}

	err = validations.ValidateExpenseObject(&expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	
	err = expense.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create expense, try again later!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Expense created!", "expense": expense})
}





func EditExpense(context *gin.Context) {
	var expense entity.Expense
	err := context.ShouldBindJSON(&expense)

	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request"})
		return
	}

	err = validations.ValidateExpenseObject(&expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}


	
	eId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request (ID parameter)"})
		return
	}

	err = expense.Modify(eId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Record updated!"})
}


func DeleteExpense(context *gin.Context){
	eId , err := strconv.ParseInt(context.Param("id"),10,64)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt parse expense ID"})
		return
	}
	
	err = entity.Delete(eId)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"Couldnt delete the record!"})
		return
	}

	context.JSON(http.StatusBadRequest, gin.H{"message":"Deleted the record!!"})

}