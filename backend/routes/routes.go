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
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	DeleteDeviceTargetTemp(c)
	DeleteBubMeasurements(c)
	DeleteTempMeasurements(c)
	c.Status(http.StatusOK)
}

// GetDeviceTargetTemps returns all devices target temps
func GetDeviceTargetTemps(c *gin.Context) {
	devices, err := database.GetAllDeviceTargetTemps(c)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, devices)
	}
}

// GetDeviceTargetTemp returns a device target temp
func GetDeviceTargetTemp(c *gin.Context) {
	deviceName := c.Param("device_name")
	device, err := database.GetDeviceTargetTemp(c, deviceName)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, device)
	}
}

// PostDeviceTargetTemp inserts a device target temp to a database
func PostDeviceTargetTemp(c *gin.Context) {
	deviceName := c.Param("device_name")
	device, err := database.GetDevice(c, deviceName)
	if err != nil || !helpers.Contains(device.Censors, "temperature") {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	var tempMeasurementRead helpers.TempMeasurementRead
	c.BindJSON(&tempMeasurementRead)
	deviceTargetTemp := database.DeviceTargetTemp{
		Device:          device.Name,
		Value:           helpers.StringToFloatConverter(tempMeasurementRead.Value),
		MeasurementUnit: tempMeasurementRead.MeasurementUnit,
	}
	err = database.UpsertDeviceTargetTemp(c, deviceTargetTemp)
	if err != nil {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}
}

// DeleteDeviceTargetTemp deletes a device target temp to a database
func DeleteDeviceTargetTemp(c *gin.Context) {
	deviceName := c.Param("device_name")
	err := database.DeleteDeviceTargetTemp(c, deviceName)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}

// GetAllBubMeasurements returns all bubble measurements
func GetAllBubMeasurements(c *gin.Context) {
	startTime, endTime := helpers.ParamReader(c.Param("start_time"),
		c.Param("end_time"))
	bubMeasurements, err := database.GetAllBubMeasurements(c,
		startTime, endTime)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, bubMeasurements)
	}
}

// GetBubMeasurements returns bubble measurements from a time frame
func GetBubMeasurements(c *gin.Context) {
	deviceName := c.Param("device_name")
	startTime, endTime := helpers.ParamReader(c.Param("start_time"),
		c.Param("end_time"))
	bubMeasurement, err := database.GetBubMeasurements(c, deviceName,
		startTime, endTime)
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
	if err != nil || !helpers.Contains(device.Censors, "bubble") {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	bubMeasurement := database.BubMeasurement{
		Device:    device.Name,
		Timestamp: time.Now().UnixNano(),
	}
	err = database.InsertBubMeasurement(c, bubMeasurement)
	if err != nil {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}

}

// DeleteBubMeasurements deletes bubble measurements from a time frame
func DeleteBubMeasurements(c *gin.Context) {
	deviceName := c.Param("device_name")
	startTime, endTime := helpers.ParamReader(c.Param("start_time"),
		c.Param("end_time"))
	err := database.DeleteBubMeasurements(c, deviceName,
		startTime, endTime)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}

// GetAllTempMeasurements returns all temp measurements
func GetAllTempMeasurements(c *gin.Context) {
	startTime, endTime := helpers.ParamReader(c.Param("start_time"),
		c.Param("end_time"))
	tempMeasurements, err := database.GetAllTempMeasurements(c,
		startTime, endTime)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, tempMeasurements)
	}
}

// GetTempMeasurements returns temp measurements from a time frame
func GetTempMeasurements(c *gin.Context) {
	deviceName := c.Param("device_name")
	startTime, endTime := helpers.ParamReader(c.Param("start_time"),
		c.Param("end_time"))
	tempMeasurement, err := database.GetTempMeasurements(c, deviceName,
		startTime, endTime)
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
	if err != nil || !helpers.Contains(device.Censors, "temperature") {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	var tempMeasurementRead helpers.TempMeasurementRead
	c.BindJSON(&tempMeasurementRead)
	tempMeasurement := database.TempMeasurement{
		Device:          device.Name,
		Timestamp:       time.Now().UnixNano(),
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

// DeleteTempMeasurements deletes temp measurements from a time frame
func DeleteTempMeasurements(c *gin.Context) {
	deviceName := c.Param("device_name")
	startTime, endTime := helpers.ParamReader(c.Param("start_time"),
		c.Param("end_time"))
	err := database.DeleteTempMeasurements(c, deviceName,
		startTime, endTime)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}
