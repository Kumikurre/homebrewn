package main

import (
	"net/http"
	
	gin "github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()

	// Gin should probably not route any static content, but a web server should do it instead...
	router.Static("/assets/", "../frontend/")

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

