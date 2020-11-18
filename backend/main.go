package main

import (
	"net/http"

	gin "github.com/gin-gonic/gin"

	"github.com/Kumikurre/homebrewn/backend/database"
)

func main() {
	database.Init("mongodb://mongo:27017")

	router := gin.Default()

	// List all devices
	router.GET("/devices", func(c *gin.Context) {
		devices := database.GetAllDevices(c)
		c.JSON(http.StatusOK, devices)
	})

	// Get more specific data about one device
	router.GET("/devices/:name", func(c *gin.Context) {
		c.String(http.StatusOK, "get device data")
	})

	router.POST("/device", func(c *gin.Context) {
		var device database.Device
		c.BindJSON(&device)
		res := database.InsertDevice(c, device)

		c.String(http.StatusOK, "device %v posted", res.InsertedID)
	})

	router.Run(":8080")
}
