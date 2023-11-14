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

        c.JSON(http.StatusCreated, account.ReturnSafeInfo())
    }
}

func UpdateAccount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var account models.Account
		if err := c.ShouldBindJSON(&account); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if result := db.Save(&account); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, account.ReturnSafeInfo())
	}
}

func DeleteAccountFromTelnumber (db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var account models.Account
		if result := db.Where("tel_number = ?", c.Param("telnumber")).Delete(&account); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
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

func GetAccountInfoFromTelnumber (db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var account models.Account
		if result := db.Where("tel_number = ?", c.Param("telnumber")).Find(&account); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, account.ReturnSafeInfo())
	}
}
