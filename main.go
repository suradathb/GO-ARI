package main

import (
	_ "go-ari/docs" // swagger documentation
	"go-ari/routes"
	"log"

	_ "go-ari/docs" // swagger documentation

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Go Swagger API Example
// @version 1.0
// @description This is a sample server for a Go Swagger API example.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	// Configure trusted proxies
	err := r.SetTrustedProxies([]string{"192.168.1.1", "192.168.1.2"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	routes.RegisterBookRoutes(r)
	routes.RegisterProductRoutes(r)
	// Start server
	r.Run(":8080")
}
