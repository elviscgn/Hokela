package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "Healthy",
		"timestamp": time.Now().Unix(),
	})
}

func SetupRouter(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	customerHandler := &CustomerHandler{DB: db}

	api := router.Group("/api/v1")
	{
		api.GET("/ping", HealthCheck)
		api.GET("/customers/:id", customerHandler.GetCustomer)
		api.POST("/customers", customerHandler.CreateCustomer)
		api.PUT("/customers/:id", customerHandler.UpdateCustomer)
		api.DELETE("/customers/:id", customerHandler.DeleteCustomer)

	}

	return router

}
