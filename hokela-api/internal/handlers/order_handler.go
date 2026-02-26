package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderHandler struct {
	DB *gorm.DB
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {

	var order Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order data provided. Please check your fields."})
		return
	}

	order.OrderID = uuid.New().String()

	if err := h.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order to the database. Please try again later."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   order})

}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	var order Order
	orderId := c.Param("id")

	if err := h.DB.Preload("Items").Where("order_id = ?", orderId).First(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order found successfully",
		"order":   order,
	})
}

func (h *OrderHandler) GetOrdersByCustomer(c *gin.Context) {
	customerId := c.Param("customer_id")
	var orders []Order
	if err := h.DB.Preload("Items").Where("customer_id = ?", customerId).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Orders for this customer not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Orders for customer found successfully",
		"orders":  orders,
	})
}

func (h *OrderHandler) GetOrdersBySpaza(c *gin.Context) {
	spazaId := c.Param("spaza_id")
	var orders []Order

	if err := h.DB.Preload("Items").Where("spaza_id = ?", spazaId).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Orders for this spaza not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Orders for spaza found successfully",
		"orders":  orders,
	})
}
