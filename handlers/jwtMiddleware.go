package handlers

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
		c.Abort()
		return
	}

	tokenStr := bearerToken[1]

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(toekn *jwt.Token) (interface{}, error) {
		return (jwtKey), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("tel_number", claims.TelNumber)

	c.Next()

}}