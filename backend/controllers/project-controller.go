package controllers

import (
	"os"
	"path/filepath"
	"strings"
	"surtr-engine/models"

	"github.com/gin-gonic/gin"
)

/*
 * [FOR MY STUPID BRAIN ON THE FUTURE]
 * ProjectController is responsible for handling project-related requests.
 *
 */

// Class ProjectController
type ProjectController struct {
	dbFolder   string
	folderName string
}

// Constructor
func NewProjectController(dbFolder string, folderName string) *ProjectController {
	projectFolder := filepath.Join(dbFolder, folderName)
	if _, err := os.Stat(projectFolder); os.IsNotExist(err) {
		os.MkdirAll(projectFolder, 0755)
	}

	return &ProjectController{dbFolder: dbFolder, folderName: folderName}
}

// Private method
func (c *ProjectController) _createFolder(projectId string, projectName string) error {
	projectNameNormalized := strings.ReplaceAll(projectName, " ", "-")
	projectFolderName := projectId + "-" + projectNameNormalized

	projectFolderCurrent := filepath.Join(c.dbFolder, c.folderName, projectFolderName)
	if _, err := os.Stat(projectFolderCurrent); os.IsNotExist(err) {
		os.MkdirAll(projectFolderCurrent, 0755)
	}
	return nil
}

// Public method
func (c *ProjectController) Create(ctx *gin.Context) {
	var project

}
