package routes

import (
	"github.com/Klambri/klambri-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Playbook(rg *gin.RouterGroup) {
	rg.GET("/", handlers.GetPlaybook)
	rg.POST("/upload", handlers.UploadPlaybook)
	rg.DELETE("/", handlers.DeletePlaybook)
	rg.PUT("/", handlers.UpdatePlaybook)
}
