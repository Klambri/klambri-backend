package routes

import (
	"github.com/Klambri/klambri-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Playbook(rg *gin.RouterGroup) {
	rg.GET("/getall", handlers.GetAllPlaybooks)
	rg.POST("/upload", handlers.UploadPlaybook)
	rg.DELETE("/delete/name/:name", handlers.DeletePlaybook)
}
