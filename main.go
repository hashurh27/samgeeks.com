// main.go

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"samgeeks/configs"
	"samgeeks/constants"
	"samgeeks/routes"
)

var user = constants.DBUsername
var password = constants.DBPassword
var host = constants.DBHost
var database = constants.DBTable

func main() {
	db, err := configs.DbConnect(user, password, host, database)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	request := gin.Default()
	routes.SetupRoutes(request)
	request.Run(":8080")
}
