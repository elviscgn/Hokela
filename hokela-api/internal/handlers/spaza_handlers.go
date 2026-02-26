package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SpazaHandler struct {
	DB *gorm.DB
}

func (h *SpazaHandler) CreateSpaza(c *gin.Context) {
	var spaza Spaza
	if err := c.ShouldBindJSON(&spaza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	spaza.SpazaID = uuid.New().String()

	if err := h.DB.Create(&spaza).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Spaza created",
		"spaza":   spaza})
}

func (h *SpazaHandler) GetAllSpazas(c *gin.Context) {
	var spazas []Spaza

	if err := h.DB.Find(&spazas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Spazas found",
		"spazas":  spazas,
	})
}

func (h *SpazaHandler) GetSpaza(c *gin.Context) {
	spazaId := c.Param("spaza_id")

	var spaza Spaza

	if err := h.DB.Where("spaza_id = ?", spazaId).First(&spaza).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error: Spaza not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Spaza found",
		"spaza":   spaza,
	})

}

func (h *SpazaHandler) UpdateSpaza(c *gin.Context) {
	spazaId := c.Param("spaza_id")
	var spaza Spaza

	if err := h.DB.Where("spaza_id = ?", spazaId).First(&spaza).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error: Spaza not found"})
		return
	}

	var updatedSpaza Spaza
	if err := c.ShouldBindJSON(&updatedSpaza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	if h.DB.Model(&spaza).Updates(updatedSpaza).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Spaza update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Spaza updated"})

}

func (h *SpazaHandler) DeleteSpaza(c *gin.Context) {
	spazaId := c.Param("spaza_id")

	if h.DB.Where("spaza_id = ?", spazaId).Delete(&Spaza{}).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error: Spaza not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Spaza deleted"})
}
