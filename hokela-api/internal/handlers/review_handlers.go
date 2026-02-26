package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewHandler struct {
	DB *gorm.DB
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var review Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data received"})
		return
	}

	review.ReviewID = uuid.New().String()
	if err := h.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Review successfully created", "review": review})
}

func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	reviewId := c.Param("id")

	var review Review

	if err := h.DB.Where("review_id = ?", reviewId).First(&review).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}
	if err := h.DB.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review successfully deleted"})

}

func (h *ReviewHandler) GetReviewsBySpaza(c *gin.Context) {
	spazaID := c.Param("id")
	var spazaReviews []Review

	if err := h.DB.Where("spaza_id = ?", spazaID).Find(&spazaReviews).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reviews for this spaza not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":       "Reviews for this spaza found",
		"spaza_reviews": spazaReviews,
	})
}

func (h *ReviewHandler) GetReviewsByRunner(c *gin.Context) {

	runnerId := c.Param("id")

	var runnerReviews []Review

	if err := h.DB.Where("runner_id = ?", runnerId).Find(&runnerReviews).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reviews for this runner not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":        "Reviews for this runner found",
		"runner_reviews": runnerReviews,
	})

}

func (h *ReviewHandler) GetReviewsByOrder(c *gin.Context) {

	orderId := c.Param("id")
	var review Review

	if err := h.DB.Where("order_id = ?", orderId).First(&review).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review for this order not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Review for this order found",
		"orderReview": review,
	})
}

func (h *ReviewHandler) GetReviewsByReviewer(c *gin.Context) {
	reviewerId := c.Param("id")
	var reviewerReviews []Review

	if err := h.DB.Where("reviewer_id = ?", reviewerId).Find(&reviewerReviews).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reviews by this reviewer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":         "Reviews by this reviewer found",
		"reviewerReviews": reviewerReviews,
	})
}
