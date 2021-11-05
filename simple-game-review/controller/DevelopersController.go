package controller

import (
	"net/http"
	"time"

	"gamereview/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DevInput struct {
	Name string `json:"name"`
}

// GetDevs godoc
// @Tags Developer
// @Produce json
// @Success 200 {object} []models.Developers
// @Router /developer/ [get]
func GetDevs(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var dev []models.Developers
	db.Find(&dev)

	c.JSON(http.StatusOK, gin.H{"data": dev})
}

// CreateDevs godoc
// @Tags Developer
// @Param Body body DevInput true "body to create data game_developers"
// @Produce json
// @Success 200 {object} models.Developers
// @Router /developer/ [post]
func CreateDevs(c *gin.Context) {
	// Validate input
	var input DevInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	dev := models.Developers{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&dev)

	c.JSON(http.StatusOK, gin.H{"data": dev})
}

// GetDevsID godoc
// @Tags Developer
// @Produce json
// @Param id path string true "Get Devs by ID"
// @Success 200 {object} []models.Developers
// @Router /developer/{id} [get]
func GetDevsID(c *gin.Context) { // Get model if exist
	var dev models.Developers

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&dev).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dev})
}

// GetGamesByDevID godoc
// @Tags Developer
// @Produce json
// @Param id path string true
// @Success 200 {object} []models.Games
// @Router /developer/{id}/games [get]
func GetGamesByDevId(c *gin.Context) { // Get model if exist
	var games []models.Games

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("dev_id = ?", c.Param("id")).Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// UpdateDev godoc
// @Tags Developer
// @Produce json
// @Param id path string true "dev id"
// @Param body body DevInput true "update the developer"
// @Success 200 {object} models.Developers
// @Router /developer/{id} [patch]
func UpdateDev(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var dev models.Developers
	if err := db.Where("id = ?", c.Param("id")).First(&dev).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input DevInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Developers
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&dev).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": dev})
}

// CreateDevs godoc
// @Tags Developer
// @Produce json
// @Param id path string true
// @Success 200 {object} models.Developers
// @Router /developer/{id} [delete]
func DeleteDev(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var dev models.Developers
	if err := db.Where("id = ?", c.Param("id")).First(&dev).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&dev)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
