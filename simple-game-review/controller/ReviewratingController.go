package controller

import (
	"net/http"
	"time"

	"gamereview/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RatingInput struct {
	ReviewsID uint `json:"review_id"`
	Helpful   bool `json:"helpful"`
}

// GetRate godoc
// @Tags ReviewRating
// @Produce json
// @Success 200 {object} []models.ReviewRating
// @Router /rating/ [get]
func GetRate(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var rate []models.ReviewRating
	db.Find(&rate)

	c.JSON(http.StatusOK, gin.H{"data": rate})
}

// CreateRate godoc
// @Tags ReviewRating
// @Produce json
// @Param Body body RatingInput true
// @Success 200 {object} []models.ReviewRating
// @Router /rating/ [post]
func CreateRate(c *gin.Context) {
	// Validate input
	var input RatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	rate := models.ReviewRating{ReviewsID: input.ReviewsID, Helpful: input.Helpful}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&rate)

	c.JSON(http.StatusOK, gin.H{"data": rate})
}

// GetRateID godoc
// @Tags ReviewRating
// @Produce json
// @Param id path string true
// @Success 200 {object} []models.ReviewRating
// @Router /rating/{id} [get]
func GetRateID(c *gin.Context) { // Get model if exist
	var rate models.ReviewRating

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&rate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rate})
}

// UpdateRate godoc
// @Tags ReviewRating
// @Produce json
// @Param id path string true
// @Param id path RatingInput true
// @Success 200 {object} []models.ReviewRating
// @Router /rating/{id} [patch]
func UpdateRate(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var rate models.ReviewRating
	if err := db.Where("id = ?", c.Param("id")).First(&rate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input RatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.ReviewRating
	updatedInput.ReviewsID = input.ReviewsID
	updatedInput.Helpful = input.Helpful
	updatedInput.UpdatedAt = time.Now()

	db.Model(&rate).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": rate})
}

// DeleteRate godoc
// @Tags ReviewRating
// @Produce json
// @Param id path string true
// @Success 200 {object} []models.ReviewRating
// @Router /rating/{id} [delete]
func DeleteRate(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var rate models.ReviewRating
	if err := db.Where("id = ?", c.Param("id")).First(&rate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&rate)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
