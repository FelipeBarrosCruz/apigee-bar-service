package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func getWebServerPort() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "80"
	}
	return fmt.Sprintf(":%s", port)
}

func main() {
	PORT := getWebServerPort()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "bar-service",
		})
	})

	r.GET("/name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "Camila",
		})
	})

	r.Run(PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
