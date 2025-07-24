package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret")

func JwtKey() []byte { return jwtKey }

type Credentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	UserID uint  `json:"user_id"`
	Exp    int64 `json:"exp"`
}

func Register(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// In a real app, save user to DB here
	c.JSON(http.StatusOK, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// In a real app, validate user from DB here
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := map[string]interface{}{
		"user_id": 1,
		"exp":     expirationTime.Unix(),
	}
	tokenString, _ := jwt.Sign(claims, jwtKey)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func ForgotPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "password reset link sent"})
}
