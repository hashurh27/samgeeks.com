package routes

import (
	"github.com/gin-gonic/gin"
	"samgeeks/api/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/user/:id", handlers.GetUserByID)
	router.POST("/user", handlers.CreateUser)
	router.POST("/post", handlers.CreatePost)
	router.GET("/post/:id", handlers.GetPostByID)
	//	router.POST("/user", handlers.)
	//	router.GET("/post/:id", handlers.GetPostByID)
}
