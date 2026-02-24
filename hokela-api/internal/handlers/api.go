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
	spazaHandler := &SpazaHandler{DB: db}
	productHandler := &ProductHandler{DB: db}
	orderHandler := &OrderHandler{DB: db}

	api := router.Group("/api/v1")
	{
		api.GET("/ping", HealthCheck)

		api.GET("/customers/:id", customerHandler.GetCustomer)
		api.GET("/customers", customerHandler.GetAllCustomers)
		api.POST("/customers", customerHandler.CreateCustomer)
		api.PUT("/customers/:id", customerHandler.UpdateCustomer)
		api.DELETE("/customers/:id", customerHandler.DeleteCustomer)

		api.GET("/spazas/:spaza_id", spazaHandler.GetSpaza)
		api.GET("/spazas", spazaHandler.GetAllSpazas)
		api.POST("/spazas", spazaHandler.CreateSpaza)
		api.PUT("/spazas/:spaza_id", spazaHandler.UpdateSpaza)
		api.DELETE("/spazas/:spaza_id", spazaHandler.DeleteSpaza)

		api.GET("/products/:id", productHandler.GetProduct)
		api.GET("/products", productHandler.GetAllProducts)
		api.GET("/spazas/:spaza_id/products", productHandler.GetProductsBySpaza)
		api.POST("/products", productHandler.AddProduct)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)

		api.GET("/orders/:id", orderHandler.GetOrder)
		api.GET("/spazas/:spaza_id/orders", orderHandler.GetOrdersBySpaza)
		api.GET("/customers/:customer_id/orders", orderHandler.GetOrdersByCustomer)
		api.POST("/orders", orderHandler.CreateOrder)
	}

	return router

}
