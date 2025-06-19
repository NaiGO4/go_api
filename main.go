package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola desde Gin!",
		})
	})

	r.GET("hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"massage": "Hola" + name + " desde Gin!",
		})
	})

	r.Run("localhost:8080") // Run on port 8080
}
