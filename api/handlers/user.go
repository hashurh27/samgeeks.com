// handlers/user_handler.go

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"samgeeks/configs"
	"samgeeks/models"
)

func GetUserByID(c *gin.Context) {
	Param := c.Param("id")
	userID, err := strconv.Atoi(Param)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid User Id"})
	}
	db, _ := configs.DbConnected()
	users, _ := models.ScanUser(db, userID)

	c.JSON(200, gin.H{"user": users[0]})
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	c.BindJSON(&newUser)
	db, _ := configs.DbConnected()

	models.InsertUser(db, newUser.Username, newUser.Email, newUser.Password)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
