package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")

	// List all devices
	router.GET("/devices", func(c *gin.Context) {
		c.String(http.StatusOK, "devices listing endpoint")
	})

	// Get more specific data about one device
	router.GET("/devices/:name/", func(c *gin.Context) {
		c.String(http.StatusOK, "get device data")
	})

	//
	router.POST("/devices/:name/", func(c *gin.Context) {
		c.String(http.StatusOK, "post device data(measurements)")
	})

	router.Run(":8080")
}

