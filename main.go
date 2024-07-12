package main

import (
	"Gin/src/config"
	"Gin/src/initializers"
	"Gin/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables() // Load Env
	config.ConnectDB()              // Connect DB

	r := gin.Default()
	routes.UserRoutes(r)
	r.Run()
}
