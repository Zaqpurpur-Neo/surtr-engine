package routes

import (
	"surtr-engine/controllers"
	"surtr-engine/models"

	"github.com/gin-gonic/gin"
)

/*
 * [FOR MY STUPID BRAIN ON THE FUTURE]
 * ApiRoutes is api router for the Surtr Engine.
 */
func ApiRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	dbFolder := "database"
	{
		apiRouter.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Welcome to the Surtr Engine API!",
			})
		})

		categoryController := controllers.NewCategoryController(dbFolder, "categories.json")
		categoryController.SeedCategories([]models.Category{
			{Name: "Digital Art", Color: "#efefef"},
			{Name: "Programming", Color: "#efefef"},
			{Name: "College Tasks", Color: "#efefef"},
		})

		apiRouter.POST("/create-category", categoryController.Create)
		apiRouter.GET("/get-categories", categoryController.GetAll)
		apiRouter.GET("/get-category/:id", categoryController.GetById)

		projectController := controllers.NewProjectController(dbFolder, "projects")
		apiRouter.POST("/create-project", projectController.Create)

	}
}
