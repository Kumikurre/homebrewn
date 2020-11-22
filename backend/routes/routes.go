package routes

import (
	"net/http"
	"time"

	gin "github.com/gin-gonic/gin"

	"github.com/Kumikurre/homebrewn/backend/database"
	"github.com/Kumikurre/homebrewn/backend/helpers"
)

// GetDevices returns all devices
func GetDevices(c *gin.Context) {
	devices, err := database.GetAllDevices(c)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, devices)
	}
}

// GetDevice returns a device
func GetDevice(c *gin.Context) {
	name := c.Param("name")
	device, err := database.GetDevice(c, name)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, device)
	}
}

// PostDevice inserts a device to a database
func PostDevice(c *gin.Context) {
	var device database.Device
	c.BindJSON(&device)
	err := database.InsertDevice(c, device)
	if err != nil {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}

}

// DeleteDevice deletes a device to a database
func DeleteDevice(c *gin.Context) {
	name := c.Param("name")
	err := database.DeleteDevice(c, name)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}

// GetBubMeasurements returns all bubble measurements
func GetBubMeasurements(c *gin.Context) {
	bubMeasurements, err := database.GetAllBubMeasurements(c)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, bubMeasurements)
	}
}

// GetBubMeasurement returns a bubble measurement
func GetBubMeasurement(c *gin.Context) {
	deviceName := c.Param("device_name")
	timestamp := helpers.StringToIntConverter(c.Param("timestamp"))
	bubMeasurement, err := database.GetBubMeasurement(c, deviceName, timestamp)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, bubMeasurement)
	}
}

// PostBubMeasurement inserts a bubble measurement to a database
func PostBubMeasurement(c *gin.Context) {
	deviceName := c.Param("device_name")
	device, err := database.GetDevice(c, deviceName)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	bubMeasurement := database.BubMeasurement{
		Device:    device,
		Timestamp: time.Now().Unix(),
	}
	err = database.InsertBubMeasurement(c, bubMeasurement)
	if err != nil {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}

}

// DeleteBubMeasurement inserts a device to a database
func DeleteBubMeasurement(c *gin.Context) {
	deviceName := c.Param("device_name")
	timestamp := helpers.StringToIntConverter(c.Param("timestamp"))
	err := database.DeleteBubMeasurement(c, deviceName, timestamp)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}

// GetTempMeasurements returns all temp measurements
func GetTempMeasurements(c *gin.Context) {
	tempMeasurements, err := database.GetAllTempMeasurements(c)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, tempMeasurements)
	}
}

// GetTempMeasurement returns a temp measurement
func GetTempMeasurement(c *gin.Context) {
	deviceName := c.Param("device_name")
	timestamp := helpers.StringToIntConverter(c.Param("timestamp"))
	tempMeasurement, err := database.GetTempMeasurement(c, deviceName, timestamp)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, tempMeasurement)
	}
}

// PostTempMeasurement inserts a temp measurement to a database
func PostTempMeasurement(c *gin.Context) {
	deviceName := c.Param("device_name")
	device, err := database.GetDevice(c, deviceName)
	var tempMeasurementRead helpers.TempMeasurementRead
	c.BindJSON(&tempMeasurementRead)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	tempMeasurement := database.TempMeasurement{
		Device:          device,
		Timestamp:       time.Now().Unix(),
		Value:           helpers.StringToFloatConverter(tempMeasurementRead.Value),
		MeasurementUnit: tempMeasurementRead.MeasurementUnit,
	}
	err = database.InsertTempMeasurement(c, tempMeasurement)
	if err != nil {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}

}

// DeleteTempMeasurement inserts a device to a database
func DeleteTempMeasurement(c *gin.Context) {
	deviceName := c.Param("device_name")
	timestamp := helpers.StringToIntConverter(c.Param("timestamp"))
	err := database.DeleteTempMeasurement(c, deviceName, timestamp)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}
