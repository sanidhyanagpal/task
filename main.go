package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fintech-app/kafka"
	"fintech-app/services"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "running"})
	})

	r.POST("/login", services.LoginHandler)
	r.POST("/register", services.RegisterHandler)

	go kafka.StartConsumer()

	r.Run(":8080")
}
