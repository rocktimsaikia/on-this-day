package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
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
