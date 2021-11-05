package controller

import (
	"net/http"
	"time"

	"gamereview/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameInput struct {
	Name         string `json:"name"`
	Year         int    `json:"year"`
	PublishersID uint   `json:"publisher_id"`
	DevelopersID uint   `json:"dev_id"`
	CategoryID   []uint `json:"category_id" gorm:"type:text[]"`
}

// GetGames godoc
// @Tags Games
// @Produce json
// @Success 200 {object} []models.Games
// @Router /games/ [get]
func GetGames(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var game []models.Games
	db.Find(&game)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// CreateGame godoc
// @Tags Games
// @Produce json
// @Param Body body GameInput true "input a game"
// @Success 200 {object} models.Games
// @Router /games/ [post]
func CreateGame(c *gin.Context) {
	// Validate input
	var input GameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	game := models.Games{Name: input.Name, Year: input.Year, PublishersID: input.PublishersID, DevelopersID: input.DevelopersID, CategoryID: input.CategoryID}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&game)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// GetGameID godoc
// @Tags Games
// @Produce json
// @Param id path string true "id game"
// @Success 200 {object} []models.Games
// @Router /games/{id} [get]
func GetGameID(c *gin.Context) { // Get model if exist
	var game models.Games

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// GetReviewByGameId godoc
// @Tags Games
// @Produce json
// @Param id path string true "id game"
// @Success 200 {object} []models.Games
// @Router /games/{id}/reviews [get]
func GetReviewByGameId(c *gin.Context) { // Get model if exist
	var rev []models.Reviews

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("game_id = ?", c.Param("id")).Find(&rev).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rev})
}

// UpdateGame godoc
// @Tags Games
// @Produce json
// @Param id path string true "id game"
// @Param Body body GameInput true "update game title/version"
// @Success 200 {object} models.Games
// @Router /games/{id} [patch]
func UpdateGame(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var game models.Games
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input GameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Games
	updatedInput.Name = input.Name
	updatedInput.Year = input.Year
	updatedInput.PublishersID = input.PublishersID
	updatedInput.DevelopersID = input.DevelopersID
	updatedInput.CategoryID = input.CategoryID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&game).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// DeleteGame godoc
// @Tags Games
// @Produce json
// @Param id path string true "id game"
// @Success 200 {object} models.Games
// @Router /games/{id} [delete]
func DeleteGame(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var game models.Games
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&game)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
