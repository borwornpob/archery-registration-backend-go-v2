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
	protected := router.Group("/").Use(handlers.JWTAuthMiddleware())

	router.POST("/login", handlers.Login((db)))

	// account routes
	router.POST("/register", handlers.RegisterAccount((db)))
	protected.PUT("/update", handlers.UpdateAccount((db)))
	router.GET("/account/:telnumber", handlers.GetAccountInfoFromTelnumber((db)))
	router.GET("/accounts", handlers.GetAllAccount((db)))
	protected.DELETE("/delete/:telnumber", handlers.DeleteAccountFromTelnumber((db)))

	

	router.Run(":8080")
}