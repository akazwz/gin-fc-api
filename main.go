package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	c := gin.Default()
	c.GET("/fc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    2000,
			"message": "success",
			"path":    "/fc",
		})
	})
	err := c.Run("0.0.0.0:9000")
	if err != nil {
		log.Println("server run error")
	}
}
