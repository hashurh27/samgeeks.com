// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"samgeeks/routes"
)

func main() {

	//db, err := configs.DbConnected()
	//if err != nil {
	//	fmt.Println("Error connecting to database:", err)
	//	return
	//}
	//models.GetPostByID(db, 7)
	//err = models.InsertPost(db, "hesam", "hesam", "slug2g", "3")
	//if err != nil {
	//	fmt.Println(err)
	//}
	route := gin.Default()
	routes.SetupRoutes(route)
	route.Run(":8080")

	// Prepare data for insertion

	//request := gin.Default()
	//routes.SetupRoutes(request)
	//request.Run(":8080")
	// Insert the post
	//postID, err := models.InsertPost(db, title, content, authorID)
	//if err != nil {
	//	fmt.Println("Error inserting post:", err)
	//	return
	//}
	//
	//fmt.Printf("Post inserted successfully with ID: %d\n", postID)
	//title := "My Awesome Post"
	//content := "This is the content of the post. It can be quite lengthy and informative."
	//authorID := "1" // Replace with the actual author ID
}

/*
var user = constants.DBUsername
var password = constants.DBPassword
var host = constants.DBHost
var database = constants.DBTable
request := gin.Default()
routes.SetupRoutes(request)
request.Run(":8080")

*/
