package main

import (
	"issue-creator/controllers"
	"issue-creator/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	db.ConnectDatabase()

	// // Routes
	r.GET("/record", controllers.FindIssues)
	// r.GET("/books/:id", controllers.FindBook)
	// r.POST("/books", controllers.CreateBook)
	// r.PATCH("/books/:id", controllers.UpdateBook)
	// r.DELETE("/books/:id", controllers.DeleteBook)

	// Run the server
	r.Run()
}
