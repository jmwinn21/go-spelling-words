package main

import (
	"encoding/json"
	"io/ioutil"
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

	router.GET("/", func(c *gin.Context) {
		jsonFile, err := os.Open("static/current.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		c.JSON(http.StatusOK, result)
	})

	router.GET("/current", func(c *gin.Context) {
		c.JSON(http.StatusOK, currentWords)
	})

	router.GET("/all", func(c *gin.Context) {
		c.JSON(http.StatusOK, currentWords)
	})

	router.Run(":" + port)
}
