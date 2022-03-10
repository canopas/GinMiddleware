package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	basicAuth := gin.BasicAuth(gin.Accounts{
		"lorem": "ipsum",
	})

	router.Use(LogEndpointURL)

	authorized := router.Group("/", basicAuth)

	authorized.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Middleware",
		})
	})

	router.Run(":8081")

}

func LogEndpointURL(c *gin.Context) {
	c.Next()
	log.Printf("Endpoint URL is %v", c.Request.URL)
}
