package handlers

import (
    "archery-registration/models"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterAccount(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var account models.Account
        if err := c.ShouldBindJSON(&account); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if result := db.Create(&account); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(http.StatusCreated, account)
    }
}

func GetAllAccount (db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var accounts []models.Account
        if result := db.Find(&accounts); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(http.StatusOK, accounts)
	}
}
