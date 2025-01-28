package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Playbook struct {
	Name string `json:"name"`
}

// GetAllPlaybooks godoc
// @Summary         Get all playbooks
// @Description     Get all playbooks
// @Produce         application/json
// @Tags            Playbook
// @Success         200 {array} Playbook
// @Router          /playbook/getall [get]
func GetAllPlaybooks(c *gin.Context) {
	files, err := os.ReadDir("./playbooks")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to reading playbooks folder"})
		return
	}

	playbooks := make([]Playbook, len(files))

	for i := 0; i < len(files); i++ {
		playbooks[i] = Playbook{Name: files[i].Name()}
	}

	fmt.Println(playbooks)

	c.JSON(http.StatusOK, playbooks)
}

// UploadPlaybook godoc
// @Summary         Upload playbook
// @Description     Upload playbook to the server.
// @Param           file formData file true "Upload playbook file"
// @Produce         application/json
// @Tags            Playbook
// @Success         200 {object} map[string]string
// @Router          /playbook/upload [post]
func UploadPlaybook(c *gin.Context) {
	err := os.MkdirAll("./playbooks", os.ModePerm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create playbooks folder"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Upload file error"})
		return
	}

	if err := c.SaveUploadedFile(file, "./playbooks/"+file.Filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Save file error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

// DeletePlaybook godoc
// @Summary         Delete playbook
// @Description     Delete playbook from the server.
// @Param           name  path    string  true "Delete playbook file"
// @Produce         application/json
// @Tags            Playbook
// @Success         200 {object} map[string]string
// @Failure         400 {object} map[string]string
// @Router          /playbook/delete/name/{name} [delete]
func DeletePlaybook(c *gin.Context) {
	path := "./playbooks/" + c.Param("name")
	if err := os.Remove(path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Remove file error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File removed successfully"})
}
