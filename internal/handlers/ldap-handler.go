package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetPlaybook(c *gin.Context) {
	c.JSON(http.StatusCreated, "get_Playbook")
}

func UploadPlaybook(c *gin.Context) {
	err := os.MkdirAll("./playbooks", os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Upload file error: %s", err.Error())
		return
	}

	if err := c.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
		c.String(http.StatusInternalServerError, "Save file error: %s", err.Error())
		return
	}

	c.String(http.StatusOK, "File %s uploaded successfully.", file.Filename)
}

func UpdatePlaybook(c *gin.Context) {
	c.JSON(http.StatusOK, "update_Playbook")
}

func DeletePlaybook(c *gin.Context) {
	c.JSON(http.StatusCreated, "delete_Playbook")
}
