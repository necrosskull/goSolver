package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"os"
	"strings"
)

type Question struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Question string `json:"question"`
	Answers  any    `json:"answers"`
}

var searchResults []Question

func init() {
	// Load the JSON data from the file only once during initialization
	data, err := os.ReadFile("./backend/questions.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var questions []Question
	err = json.Unmarshal(data, &questions)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	fmt.Println("Loaded", len(questions), "questions")

	searchResults = questions
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.Static("/", "./static")

	router.POST("/questions", func(c *gin.Context) {
		var searchQuery struct {
			SearchQuery string `json:"searchQuery"`
		}

		if err := c.BindJSON(&searchQuery); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		searchQueryLower := strings.ToLower(strings.TrimSpace(searchQuery.SearchQuery))
		results := make([]Question, 0)

		for _, question := range searchResults {
			if strings.HasPrefix(strings.ToLower(question.Question), searchQueryLower) {
				results = append(results, question)
			}
		}

		if len(results) > 5 {
			results = results[:5]
		}

		c.JSON(200, results)
	})

	router.Run(":8082")
}
