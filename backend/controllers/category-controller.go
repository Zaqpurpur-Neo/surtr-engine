package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"surtr-engine/models"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

/*
 * [FOR MY STUPID BRAIN ON THE FUTURE]
 * CategoryController is responsible for handling category-related requests.
 */

// Class CategoryController
type CategoryController struct {
	dbFolder string
	fileName string
}

// Constructor
func NewCategoryController(dbFolder, fileName string) *CategoryController {
	categoryFile := filepath.Join(dbFolder, fileName)
	if _, err := os.Stat(categoryFile); os.IsNotExist(err) {
		os.Create(categoryFile)
	}

	categoryController := &CategoryController{
		dbFolder: dbFolder,
		fileName: fileName,
	}
	categoryController._fillEmptyJson()
	return categoryController
}

// Private method

func (c *CategoryController) _fillEmptyJson() {
	categoryFile := filepath.Join(c.dbFolder, c.fileName)
	if _, err := os.Stat(categoryFile); os.IsNotExist(err) {
		os.Create(categoryFile)
	}

	contentJson, err := json.Marshal([]models.Category{})
	if err != nil {
		return
	}
	os.WriteFile(categoryFile, contentJson, 0644)
}

func (c *CategoryController) _readCategories() ([]models.Category, error) {
	categoryFile := filepath.Join(c.dbFolder, c.fileName)
	contentJson, err := os.ReadFile(categoryFile)
	if err != nil {
		return nil, err
	}

	var categories []models.Category
	if err := json.Unmarshal(contentJson, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *CategoryController) _writeCategories(categories []models.Category) error {
	categoryFile := filepath.Join(c.dbFolder, c.fileName)
	contentJson, err := json.Marshal(categories)
	if err != nil {
		return err
	}

	return os.WriteFile(categoryFile, contentJson, 0644)
}

// Public method

func (c *CategoryController) SeedCategories(categories []models.Category) {

	// add nanoid please
	for i := range categories {
		categories[i].Id, _ = gonanoid.New(8)
	}

	if err := c._writeCategories(categories); err != nil {
		fmt.Println(err)
	}
}

func (c *CategoryController) Create(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	category.Id, _ = gonanoid.New(8)

	categories, err := c._readCategories()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	categories = append(categories, category)
	if err := c._writeCategories(categories); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, category)
}

func (c *CategoryController) GetById(ctx *gin.Context) {
	categoryId := ctx.Param("id")

	categories, err := c._readCategories()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, category := range categories {
		if category.Id == categoryId {
			ctx.JSON(200, category)
			return
		}
	}

	ctx.JSON(404, gin.H{"error": "category not found"})
}

func (c *CategoryController) GetAll(ctx *gin.Context) {
	categories, err := c._readCategories()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, categories)
}
