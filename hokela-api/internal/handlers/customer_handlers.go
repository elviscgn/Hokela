package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	DB *gorm.DB
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	customer.CustomerID = uuid.New().String()

	if err := h.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create customer. Email might be taken."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Customer created successfully!",
		"customer": customer,
	})
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	idStr := c.Param("id")
	var customer Customer

	if h.DB.Where("customer_id = ?", idStr).First(&customer).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Customer found successfully!",
		"customer": customer,
	})

}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	idStr := c.Param("id")
	var customer Customer
	if h.DB.Where("customer_id = ?", idStr).First(&customer).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	var updatedCustomer Customer
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	if h.DB.Model(&customer).Updates(updatedCustomer).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer."})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully!"})

}
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	idStr := c.Param("id")
	if h.DB.Where("customer_id = ?", idStr).Delete(&Customer{}).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully!"})
}
