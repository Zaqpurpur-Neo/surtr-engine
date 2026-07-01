package main

import (
	"surtr-engine/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.ApiRoutes(router)

	router.Run(":4455")
	print("Server running at port 4455")
}
