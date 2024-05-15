// controllers/product_controller.go
package controllers

import (
	"go-ari/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all products
// @Description Get details of all products
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, models.Products)
}

// @Summary Get a product
// @Description Get details of a product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	for _, product := range models.Products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product to create"
// @Success 201 {object} models.Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.Products = append(models.Products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

// @Summary Update a product
// @Description Update a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body models.Product true "Product to update"
// @Success 200 {object} models.Product
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	for i, product := range models.Products {
		if product.ID == id {
			if err := c.ShouldBindJSON(&product); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			models.Products[i] = product
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 {string} string "No Content"
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, product := range models.Products {
		if product.ID == id {
			models.Products = append(models.Products[:i], models.Products[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
