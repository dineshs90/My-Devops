package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	router := gin.Default()
	router.GET("/home", func(c *gin.Context) {
		c.String(http.StatusAccepted, "Welcome to Golang !!!")
	})
	//health checks
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusAccepted, "Application status: Healthy ✅")
	})
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusAccepted, "Everything is working fine !!!")
	})
	// 👉 Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Run(":8080")

}
