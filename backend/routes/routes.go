package routes

import (
	"net/http"

	gin "github.com/gin-gonic/gin"

	"github.com/Kumikurre/homebrewn/backend/database"
)

// GetDevices returns all devices
func GetDevices(c *gin.Context) {
	devices := database.GetAllDevices(c)
	c.JSON(http.StatusOK, devices)
}

// GetDevice returns a device
func GetDevice(c *gin.Context) {
	c.String(http.StatusOK, "get device data")
}

// PostDevice inserts a device to a database
func PostDevice(c *gin.Context) {
	var device database.Device
	c.BindJSON(&device)
	res := database.InsertDevice(c, device)

	c.String(http.StatusOK, "device %v posted", res.InsertedID)
}
