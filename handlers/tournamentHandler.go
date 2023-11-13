package handlers

import (
	"archery-registration/models"
	"net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func RegisterTournament(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tournament models.Tournament
		if err := c.ShouldBindJSON(&tournament); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Create(&tournament); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, tournament)
	}
}

func GetTournamentFromId(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tournament models.Tournament
		if result := db.Where("id = ?", c.Param("id")).Find(&tournament); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, tournament)
	}
}

func GetAllTournament(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tournaments []models.Tournament
		if result := db.Find(&tournaments); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, tournaments)
	}
}

func ModifyTournament(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tournament models.Tournament
		if err := c.ShouldBindJSON(&tournament); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Save(&tournament); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, tournament)
	}
}

func DeleteTournament(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tournament models.Tournament
		if err := c.ShouldBindJSON(&tournament); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Delete(&tournament); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, tournament)
	}
}