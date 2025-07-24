package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatCS(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "Hello, how can I help you?"})
}

func ChatPsychologist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "I'm here to listen. Tell me more."})
}
