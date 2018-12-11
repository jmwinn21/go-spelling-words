package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type WordsResponse struct {
	Words []string `json:"words"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	currentWords := WordsResponse{
		Words: []string{
			"triathlon",
			"trilogy",
			"trimester",
			"trident",
			"quadrangle",
			"quartet",
			"quart",
			"pentagon",
			"pentagram",
			"quintuplet",
		},
	}

	router.GET("/current", func(c *gin.Context) {
		c.JSON(http.StatusOK, currentWords)
	})

	router.Run(":" + port)
}
