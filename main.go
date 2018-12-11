package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
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

	router.GET("/", func(c *gin.Context) {
		jsonFile, err := os.Open("static/current.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var result WordsResponse
		var output WordsResponse
		err = json.Unmarshal(byteValue, &result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		output.Words = Shuffle(result.Words)

		c.JSON(http.StatusOK, output)
	})

	router.GET("/all", func(c *gin.Context) {
		jsonFile, err := os.Open("static/all.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var result WordsResponse
		var output WordsResponse
		err = json.Unmarshal(byteValue, &result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		output.Words = ShuffleSize(result.Words, 10)

		c.JSON(http.StatusOK, output)
	})

	router.Run(":" + port)
}

func Shuffle(vals []string) []string {
	dest := make([]string, len(vals))
	perm := rand.Perm(len(vals))
	for i, v := range perm {
		dest[v] = vals[i]
	}
	return dest
}

func ShuffleSize(vals []string, size int) []string {
	dest := make([]string, len(vals))
	perm := rand.Perm(len(vals))
	for i, v := range perm {
		dest[v] = vals[i]
	}
	return dest[0:size]
}
