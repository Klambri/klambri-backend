package routes

import (
	"github.com/Klambri/klambri-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func userRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")
	{
		userGroup.GET("/", handlers.GetUsers)
		userGroup.POST("/", handlers.DeleteUser)
		userGroup.DELETE("/", handlers.DeleteUser)
		userGroup.PUT("/", handlers.DeleteUser)
	}
}
