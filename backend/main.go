package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
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

func showResults(c *gin.Context) {
	searchQuery := c.Param("query")

	searchQueryLower := strings.ToLower(strings.TrimSpace(searchQuery))
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

}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./static", false)))

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/search/:query", showResults)
	}

	router.Run(":8082")
}
