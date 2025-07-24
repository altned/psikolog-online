package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Article struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var articles = []Article{
	{ID: 1, Title: "Mental Health 101", Body: "Be kind to yourself."},
	{ID: 2, Title: "Stress Management", Body: "Take a deep breath."},
}

func ListArticles(c *gin.Context) {
	c.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, a := range articles {
		if int(a.ID) == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}
