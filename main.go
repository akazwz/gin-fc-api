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
	s := &http.Server{
		Addr:    "0.0.0.0:9000",
		Handler: c,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("listenAndServe: %+v", err)
	}
}
