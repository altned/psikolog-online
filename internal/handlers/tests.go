package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Test struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

var tests = []Test{
	{ID: 1, Name: "MBTI"},
	{ID: 2, Name: "Stress"},
}

func ListTests(c *gin.Context) {
	c.JSON(http.StatusOK, tests)
}

func StartTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "test started"})
}

func SubmitTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "answers submitted"})
}

func GetTestResult(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"id": id, "result": "This is your result"})
}
