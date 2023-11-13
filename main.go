package main

import (
	"archery-registration/handlers"
	"archery-registration/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := connectDB()

	db.AutoMigrate(&models.Account{}, &models.Entry{}, &models.Tournament{})

	router := gin.Default()

	router.POST("/register", handlers.RegisterAccount((db)))
	router.GET("/accounts", handlers.GetAllAccount((db)))

	router.Run(":8080")

}