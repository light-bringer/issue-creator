package main

import (
	"fmt"
	"issue-creator/controllers"
	"issue-creator/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xanzy/go-gitlab"
)

func main() {
	r := gin.Default()

	// Connect to database
	db.ConnectDatabase()
	git, err := gitlab.NewClient("kFHvzdQTRqwyyvJHbzCg", gitlab.WithBaseURL("http://10.21.58.72/rilservicesgitlab//api/v4"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	users, _, err := git.Issues.ListProjectIssues(124, &gitlab.ListProjectIssuesOptions{})
	fmt.Printf("length of issues : %d\n", (len(users)))
	//log.Fatalf(err.Error())
	// for i, s := range users {
	// 	//fmt.Println(i, s)
	// 	//fmt.Println(len(users))
	// }

	controllers.CreateRecord("123", "324")

	// // Routes
	r.GET("/record", controllers.FindIssues)
	// r.GET("/books/:id", controllers.FindBook)
	// r.POST("/books", controllers.CreateBook)
	// r.PATCH("/books/:id", controllers.UpdateBook)
	// r.DELETE("/books/:id", controllers.DeleteBook)

	// Run the server
	r.Run()
}
