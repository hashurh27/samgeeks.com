package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"samgeeks/configs"
	"samgeeks/models"
	"strconv"
	"strings"
)

func GetPostByID(c *gin.Context) {
	Param := c.Param("id")
	UserID, err := strconv.Atoi(Param)
	if err != nil {
		c.JSON(400, gin.H{"err USER ID INVALID": "code 200"})
		return // Exit early if conversion fails
	}

	db, _ := configs.DbConnected()

	post, err := models.GetPostByID(db, UserID)
	if err != nil {
		// Handle error appropriately, considering specific errors and logging
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post}) // Directly use the post object
}

func CreatePost(c *gin.Context) {
	var newPost models.Post

	// Bind request body to newPost and handle binding errors
	if err := c.BindJSON(&newPost); err != nil {
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

	// Insert post and handle potential errors
	err = models.InsertPost(db, newPost.Title, newPost.Content, newPost.Slug, newPost.AuthorID)
	if err != nil {
		// Handle specific errors as before
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": "post with same slug already exists"})
		} else {
			// Log more detailed error information
			log.Printf("error creating post: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	// Successful post creation, handle as needed
	c.JSON(http.StatusCreated, gin.H{"message": "post created successfully"})
}
