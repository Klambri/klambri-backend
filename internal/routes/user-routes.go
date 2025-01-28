package routes

import (
	"github.com/Klambri/klambri-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func User(rg *gin.RouterGroup) {
	rg.GET("/", handlers.GetUsers)
	rg.POST("/", handlers.DeleteUser)
	rg.DELETE("/", handlers.DeleteUser)
	rg.PUT("/", handlers.DeleteUser)
}
