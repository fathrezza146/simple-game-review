package controller

import (
	"net/http"
	"time"

	"gamereview/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PubInput struct {
	Name string `json:"name"`
}

// GetPub godoc
// @Tags Publisher
// @Produce json
// @Success 200 {object} []models.Publishers
// @Router /publisher/ [get]
func GetPub(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var pub []models.Publishers
	db.Find(&pub)

	c.JSON(http.StatusOK, gin.H{"data": pub})
}

// CreatePubs godoc
// @Tags Publisher
// @Produce json
// @Param Body body PubInput true
// @Success 200 {object} []models.Publishers
// @Router /publisher/ [post]
func CreatePubs(c *gin.Context) {
	// Validate input
	var input PubInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	pub := models.Publishers{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&pub)

	c.JSON(http.StatusOK, gin.H{"data": pub})
}

// GetPubID godoc
// @Tags Publisher
// @Produce json
// @Param id path string true
// @Success 200 {object} []models.Publishers
// @Router /publisher/{id} [get]
func GetPubID(c *gin.Context) { // Get model if exist
	var pub models.Publishers

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&pub).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pub})
}

// GetGamesByPubId godoc
// @Tags Publisher
// @Produce json
// @Param id path string true
// @Success 200 {object} []models.Publishers
// @Router /publisher/{id}/games [get]
func GetGamesByPubId(c *gin.Context) { // Get model if exist
	var games []models.Games

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("publisher_id = ?", c.Param("id")).Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// UpdatePub godoc
// @Tags Publisher
// @Produce json
// @Param id path string true
// @Param id path PubInput true
// @Success 200 {object} []models.Publishers
// @Router /publisher/{id} [patch]
func UpdatePub(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var pub models.Publishers
	if err := db.Where("id = ?", c.Param("id")).First(&pub).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input PubInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Publishers
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&pub).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": pub})
}

// DeletePub godoc
// @Tags Publisher
// @Produce json
// @Param id path string true
// @Success 200 {object} []models.Publishers
// @Router /publisher/{id} [delete]
func DeletePub(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var pub models.Publishers
	if err := db.Where("id = ?", c.Param("id")).First(&pub).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&pub)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
