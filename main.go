package main

import (
	"github.com/gin-gonic/gin"

	"shanker.com/expense-tracker/controller"
	"shanker.com/expense-tracker/db"
)

func main(){
	server := gin.Default()

	db.InitDB()
	
	controller.RegisterRoutes(server)

	server.Run(":8081")
}