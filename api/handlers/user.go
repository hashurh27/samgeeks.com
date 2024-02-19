// handlers/user_handler.go

package handlers

import (
	"log"
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

	// Bind request body to newUser and handle binding errors
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Connect to database with error handling
	db, err := configs.DbConnected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to database"})
		return
	}
	defer db.Close() // Ensure database connection is closed

	// Insert user and handle potential errors
	err = models.InsertUser(db, newUser.Username, newUser.Email, newUser.Password)
	if err != nil {
		// Handle specific errors as before
		if err == models.ErrDuplicateUsername {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		} else if err == models.ErrDuplicateEmail {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		} else {
			// Log more detailed error information
			log.Printf("error creating user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	// Successful user creation, handle as needed
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}
