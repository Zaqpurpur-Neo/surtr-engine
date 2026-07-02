package controllers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"surtr-engine/models"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
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
func (c *ProjectController) _createFolder(projectId string) (string, error) {
	projectFolderCurrent := filepath.Join(c.dbFolder, c.folderName, projectId)
	if _, err := os.Stat(projectFolderCurrent); os.IsNotExist(err) {
		os.MkdirAll(projectFolderCurrent, 0755)
	}

	return filepath.Join(projectFolderCurrent, "project.json"), nil
}

func (c *ProjectController) _createProjectJSON(project models.Project, projectPath string) error {
	projectJSON, err := json.Marshal(project)
	if err != nil {
		return err
	}
	return os.WriteFile(projectPath, projectJSON, 0644)
}

func (c *ProjectController) _getProjectEntries() []os.DirEntry {
	entries, err := os.ReadDir(filepath.Join(c.dbFolder, c.folderName))
	if err != nil {
		return nil
	}
	return entries
}

// Public method

func (c *ProjectController) Create(ctx *gin.Context) {
	var projectItem models.Project
	if err := ctx.ShouldBindJSON(&projectItem); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if projectItem.ID == "" {
		projectItem.ID, _ = gonanoid.New(8)
	}
	projectItem.CreatedAt = time.Now().Format(time.RFC3339)
	projectItem.UpdatedAt = projectItem.CreatedAt

	projectPath, err := c._createFolder(projectItem.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := c._createProjectJSON(projectItem, projectPath); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Project created successfully"})
}

func (c *ProjectController) GetByID(ctx *gin.Context) {
	projectFolder := filepath.Join(c.dbFolder, c.folderName, ctx.Param("id"))
	if _, err := os.Stat(projectFolder); os.IsNotExist(err) {
		ctx.JSON(404, gin.H{"error": "Project not found"})
		return
	}

	projectPath := filepath.Join(projectFolder, "project.json")
	projectJSON, err := os.ReadFile(projectPath)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var project models.Project
	if err := json.Unmarshal(projectJSON, &project); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, project)
}

func (c *ProjectController) GetAll(ctx *gin.Context) {
	projects := []models.Project{}
	for _, entry := range c._getProjectEntries() {
		projectFolder := filepath.Join(c.dbFolder, c.folderName, entry.Name())
		projectPath := filepath.Join(projectFolder, "project.json")
		projectJSON, err := os.ReadFile(projectPath)
		if err != nil {
			continue
		}
		var project models.Project
		if err := json.Unmarshal(projectJSON, &project); err != nil {
			continue
		}
		projects = append(projects, project)
	}
	ctx.JSON(200, projects)
}

func (c *ProjectController) DeleteByID(ctx *gin.Context) {
	projectFolder := filepath.Join(c.dbFolder, c.folderName, ctx.Param("id"))
	if _, err := os.Stat(projectFolder); os.IsNotExist(err) {
		ctx.JSON(404, gin.H{"error": "Project not found"})
		return
	}

	projectPath := filepath.Join(projectFolder, "project.json")
	if err := os.Remove(projectPath); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Project deleted successfully"})
}
