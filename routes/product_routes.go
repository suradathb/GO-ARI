// routes/product_routes.go
package routes

import (
	"go-ari/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/products", controllers.GetProducts)
		api.GET("/products/:id", controllers.GetProduct)
		api.POST("/products", controllers.CreateProduct)
		api.PUT("/products/:id", controllers.UpdateProduct)
		api.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
