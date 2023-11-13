package main

import (
	"archery-registration/handlers"
	"archery-registration/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	db := connectDB()

	db.AutoMigrate(&models.Account{}, &models.Entry{}, &models.Tournament{})

	router := gin.Default()

	// account routes
	router.POST("/register", handlers.RegisterAccount((db)))
	router.PUT("/update", handlers.UpdateAccount((db)))
	router.GET("/account/:telnumber", handlers.GetAccountInfoFromTelnumber((db)))
	router.GET("/accounts", handlers.GetAllAccount((db)))
	router.DELETE("/delete/:telnumber", handlers.DeleteAccountFromTelnumber((db)))

	

	router.Run(":8080")
}