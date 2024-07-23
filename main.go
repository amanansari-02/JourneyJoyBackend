package main

import (
	"JourneyJoyBackend/src/config"
	"JourneyJoyBackend/src/initializers"
	"JourneyJoyBackend/src/middleware"
	"JourneyJoyBackend/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables() // Load Env
	config.ConnectDB()              // Connect DB
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "done",
		})
	})
	routes.UserRoutes(r)
	r.Run()
}
