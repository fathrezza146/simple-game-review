package controller

import (
	"net/http"
	"time"

	"gamereview/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewInput struct {
	GamesID    uint   `json:"game_id"`
	TextReview string `json:"review" gorm:"type:text"`
}

// GetReview godoc
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Reviews
// @Router /review/ [get]
func GetReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var rev []models.Reviews
	db.Find(&rev)

	c.JSON(http.StatusOK, gin.H{"data": rev})
}

// CreateReview godoc
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Reviews
// @Router /review/ [post]
func CreateReview(c *gin.Context) {
	// Validate input
	var input ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	rev := models.Reviews{GamesID: input.GamesID, TextReview: input.TextReview}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&rev)

	c.JSON(http.StatusOK, gin.H{"data": rev})
}

// GetReviewID godoc
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Reviews
// @Router /review/{id} [get]
func GetReviewID(c *gin.Context) { // Get model if exist
	var rev models.Reviews

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&rev).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rev})
}

// GetRatingByReviewId godoc
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Reviews
// @Router /review/{id}/ratings [get]
func GetRatingByReviewId(c *gin.Context) { // Get model if exist
	var rate []models.ReviewRating

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("review_id = ?", c.Param("id")).Find(&rate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rate})
}

// UpdateReview godoc
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Reviews
// @Router /review/{id} [patch]
func UpdateReview(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var rev models.Reviews
	if err := db.Where("id = ?", c.Param("id")).First(&rev).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Reviews
	updatedInput.GamesID = input.GamesID
	updatedInput.TextReview = input.TextReview
	updatedInput.UpdatedAt = time.Now()

	db.Model(&rev).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": rev})
}

// DeleteReview godoc
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Reviews
// @Router /review/{id} [delete]
func DeleteReview(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var rev models.Reviews
	if err := db.Where("id = ?", c.Param("id")).First(&rev).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&rev)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
