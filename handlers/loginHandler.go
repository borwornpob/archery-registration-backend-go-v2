package handlers

import (
	"archery-registration/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Credentials struct {
	TelNumber string `json:"tel_number"`
	Password string `json:"password"`
}

type Claims struct {
	TelNumber string `json:"tel_number"`
	jwt.StandardClaims
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds Credentials
		var user models.Account
		jwtKey := []byte(os.Getenv("JWT_KEY"))

		// Bind the JSON body to the credentials struct
		if err := c.BindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := db.Where("tel_number = ?", creds.TelNumber).Find(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		if user.Password != creds.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
			return
		}

		expirationTime := time.Now().Add(1 * time.Hour)
		claims := &Claims{
			TelNumber: creds.TelNumber,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		fmt.Println(jwtKey)

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

}



