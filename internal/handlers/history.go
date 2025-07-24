package handlers

import "net/http"

import "github.com/gin-gonic/gin"

type HistoryItem struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func GetHistory(c *gin.Context) {
	history := []HistoryItem{
		{Type: "chat", Data: "User talked to CS"},
		{Type: "test", Data: "User took MBTI"},
	}
	c.JSON(http.StatusOK, history)
}
