package handlers

import (
    "archery-registration/models"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterEntry(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var entry models.Entry
		if err := c.ShouldBindJSON(&entry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Create(&entry); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, entry)
	}
}

func GetEntriesFromAccountId(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var entries []models.Entry
		if result := db.Where("account_id = ?", c.Param("id")).Find(&entries); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, entries)
	}
}

func GetAllEntriesWithinTournament(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var entries []models.Entry
		if result := db.Where("tournament_id = ?", c.Param("id")).Find(&entries); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, entries)
	}
}

func ModifyEntryUsingId (db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var entry models.Entry
		if err := c.ShouldBindJSON(&entry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Save(&entry); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, entry)
	}
}

func DeleteEntry(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var entry models.Entry
		if err := c.ShouldBindJSON(&entry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Delete(&entry); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, entry)
	}
}
