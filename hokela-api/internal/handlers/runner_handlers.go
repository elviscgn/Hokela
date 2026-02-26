package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RunnerHandler struct {
	DB *gorm.DB
}

func (h *RunnerHandler) CreateRunner(c *gin.Context) {
	var runner Runner

	if err := c.BindJSON(&runner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data received"})
		return
	}
	runner.RunnerID = uuid.New().String()

	if err := h.DB.Create(&runner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create Runner"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"runner": runner.RunnerID,
		"message": "Runner created successfully"})
}

func (h *RunnerHandler) UpdateRunner(c *gin.Context) {
	var runner Runner
	runnerId := c.Param("id")
	if err := h.DB.Where("runner_id = ?", runnerId).First(&runner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Runner not found"})
		return
	}

	var updatedRunner Runner

	if err := c.BindJSON(&updatedRunner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update request"})
		return
	}
	if err := h.DB.Model(&runner).Updates(updatedRunner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not update Runner"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Runner updated successfully!",
		"runner": runner})
}

func (h *RunnerHandler) DeleteRunner(c *gin.Context) {
	runnerId := c.Param("id")

	if err := h.DB.Where("runner_id = ?", runnerId).Delete(&Runner{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not delete Runner"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Runner deleted successfully!"})
}

func (h *RunnerHandler) GetRunner(c *gin.Context) {
	var runner Runner
	runnerId := c.Param("id")
	if err := h.DB.Where("runner_id = ?", runnerId).First(&runner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Runner not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Runner found successfully!",
		"runner": runner})
}

func (h *RunnerHandler) GetAllRunners(c *gin.Context) {
	var runner []Runner
	if err := h.DB.Find(&runner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not retrieve all runners"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Runners successfully retrieved!",
		"runners": runner})
}

func (h *RunnerHandler) GetAvailableRunners(c *gin.Context) {
	var availableRunners []Runner

	if err := h.DB.Where("is_online = ?", true).Find(&availableRunners).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not retrieve online runners"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Found Runner available!",
		"available_runners": availableRunners})

}
