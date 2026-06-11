package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// 1. Initialize router
	r := gin.Default()

	// 2. Initialize routes
	initializeRoutes(r)

	// 3. Run the server
	r.Run(":8080")
}
