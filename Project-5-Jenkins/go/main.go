package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/home", func(c *gin.Context) {
		c.String(http.StatusAccepted, "Welcome to Golang !!!")
	})
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusAccepted, "Application status: Healthy ✅")
	})
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusAccepted, "Everything is working fine !!!")
	})

	router.Run(":8080")

}
