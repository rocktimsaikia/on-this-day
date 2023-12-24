package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"on-this-day/scripts"
)

func main() {
	fmt.Println("Hello, World!")

	router := gin.Default()

	router.GET("/api/v1/today", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
