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
	runnerHandler := &RunnerHandler{DB: db}
	reviewHandler := &ReviewHandler{DB: db}

	api := router.Group("/api/v1")
	{
		api.GET("/ping", HealthCheck)

		//customers
		api.GET("/customers/:id", customerHandler.GetCustomer)
		api.GET("/customers", customerHandler.GetAllCustomers)
		api.POST("/customers", customerHandler.CreateCustomer)
		api.PUT("/customers/:id", customerHandler.UpdateCustomer)
		api.DELETE("/customers/:id", customerHandler.DeleteCustomer)

		//spazas
		api.GET("/spazas/:spaza_id", spazaHandler.GetSpaza)
		api.GET("/spazas", spazaHandler.GetAllSpazas)
		api.POST("/spazas", spazaHandler.CreateSpaza)
		api.PUT("/spazas/:spaza_id", spazaHandler.UpdateSpaza)
		api.DELETE("/spazas/:spaza_id", spazaHandler.DeleteSpaza)

		//products
		api.GET("/products/:id", productHandler.GetProduct)
		api.GET("/products", productHandler.GetAllProducts)
		api.GET("/spazas/:spaza_id/products", productHandler.GetProductsBySpaza)
		api.POST("/products", productHandler.AddProduct)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)

		//orders
		api.GET("/orders/:id", orderHandler.GetOrder)
		api.GET("/spazas/:spaza_id/orders", orderHandler.GetOrdersBySpaza)
		api.GET("/customers/:customer_id/orders", orderHandler.GetOrdersByCustomer)
		api.POST("/orders", orderHandler.CreateOrder)

		//runners
		api.GET("/runners", runnerHandler.GetAllRunners)
		api.GET("/available-runners", runnerHandler.GetAvailableRunners)
		api.GET("/runners/:id", runnerHandler.GetRunner)
		api.POST("/runners", runnerHandler.CreateRunner)
		api.PUT("/runners/:id", runnerHandler.UpdateRunner)
		api.DELETE("/runners/:id", runnerHandler.DeleteRunner)

		//reviews
		api.GET("/spazas/:id/reviews", reviewHandler.GetReviewsBySpaza)
		api.GET("/runners/:id/reviews", reviewHandler.GetReviewsByRunner)
		api.GET("/orders/:id/reviews", reviewHandler.GetReviewsByOrder)
		api.GET("/customers/:id/reviews", reviewHandler.GetReviewsByReviewer)
		api.POST("/reviews", reviewHandler.CreateReview)
		api.DELETE("/reviews/:id", reviewHandler.DeleteReview)
		
	}

	return router

}
