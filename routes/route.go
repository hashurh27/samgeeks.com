package routes

import (
	"github.com/gin-gonic/gin"
	"samgeeks/api/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/user/:id", handlers.GetUserByID)
	router.POST("/user", handlers.CreateUser)
}
