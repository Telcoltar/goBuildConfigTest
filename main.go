package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type greeting struct {
	ID      string `json:"id"`
	Content string `json:"title"`
}

var greetings = []greeting{
	{ID: "1", Content: "Hello"},
	{ID: "2", Content: "World!"},
}

func getGreetings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, greetings)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getGreetings)

	router.Run(":8080")
}
