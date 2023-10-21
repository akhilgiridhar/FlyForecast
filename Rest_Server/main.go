package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPrediction(data Input) Ouput {
	// Make HTTP request to  ML model API and pass in the data
	// Parse the response
	// Return the response
}

func main() {
	r := gin.Default()

	r.POST("/predict", func(c *gin.Context) {
		var input Input
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		output := getPrediction(input)
		c.JSON(http.StatusOK, output)
	})

	r.Run()
}
