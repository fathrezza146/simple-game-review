package controller

import (
	"net/http"
	"time"

	"gamereview/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryInput struct {
	Name string `json:"name"`
}

// GetCategory godoc
// @Tags Category
// @Produce json
// @Success 200 {object} []models.Category
// @Router /category [get]
func GetCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var category []models.Category
	db.Find(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetCategory godoc
// @Tags Category
// @Produce json
// @Param Body body CategoryInput true "Create a Game Category"
// @Success 200 {object} models.Category
// @Router /category/ [post]
func CreateCategory(c *gin.Context) {
	// Validate input
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	category := models.Category{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetCategoryID godoc
// @Tags Category
// @Produce json
// @Param id path string true "id category"
// @Success 200 {object} []models.Category
// @Router /category/{id} [get]
func GetCategoryID(c *gin.Context) { // Get model if exist
	var category models.Category

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetGamesByCategoryId godoc
// @Tags Category
// @Produce json
// @Param id path string true "id category"
// @Success 200 {object} []models.Games
// @Router /category/{id}/games [get]
func GetGamesByCategoryId(c *gin.Context) { // Get model if exist
	var games []models.Games

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("category_id = ?", c.Param("id")).Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// UpdateCategory godoc
// @Tags Category
// @Produce json
// @Param id path string true "id category"
// @Param Body body CategoryInput true "update the category"
// @Success 200 {object} models.Category
// @Router /category/{id} [patch]
func UpdateCategory(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Category
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&category).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory godoc
// @Tags Category
// @Produce json
// @Param id path string true "Delete a category"
// @Success 200 {object} models.Cate gory
// @Router /category/{id} [delete]
func DeleteCategory(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
