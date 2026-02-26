package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func (h *ProductHandler) AddProduct(c *gin.Context) {
	var product Product

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data provided. Please check your fields."})
		return
	}

	product.ProductID = uuid.New().String()

	if err := h.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product to the database. Please try again later."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product added",
		"product": product})
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	var products []Product

	if err := h.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products. Please try again later."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "All Products retrieved",
		"products": products,
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	var product Product
	productID := c.Param("id")

	if err := h.DB.Where("product_id = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product retrieved",
		"product": product,
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var product Product
	productID := c.Param("id")

	if err := h.DB.Where("product_id = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found."})
		return
	}

	var updatedProduct Product

	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data provided. Please check your fields."})
		return
	}

	if h.DB.Model(&product).Updates(updatedProduct).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the product. Please try again later."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated",
		"product": product,
	})

}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {

	productID := c.Param("id")

	if err := h.DB.Where("product_id = ?", productID).Delete(&Product{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the product. Please try again later."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
}

func (h *ProductHandler) GetProductsBySpaza(c *gin.Context) {

	spazaId := c.Param("spaza_id")

	var products []Product

	if err := h.DB.Where("spaza_id = ?", spazaId).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products for this Spaza. Please try again later."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Products for this spaza retrieved",
		"products": products,
	})
}
