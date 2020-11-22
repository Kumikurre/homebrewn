package main

import (
	gin "github.com/gin-gonic/gin"

	"github.com/Kumikurre/homebrewn/backend/database"
	"github.com/Kumikurre/homebrewn/backend/routes"
)

func main() {
	database.Init("mongodb://mongo:27017")
	router := gin.Default()

	router.GET("/devices", routes.GetDevices)
	router.GET("/device/:name", routes.GetDevice)
	router.POST("/device", routes.PostDevice)
	router.DELETE("/device/:name", routes.DeleteDevice)

	router.GET("/bub_measurements", routes.GetBubMeasurements)
	router.GET("/bub_measurement/:device_name/:timestamp", routes.GetBubMeasurement)
	router.POST("/bub_measurement/:device_name", routes.PostBubMeasurement)
	router.DELETE("/bub_measurement/:device_name/:timestamp", routes.DeleteBubMeasurement)

	router.GET("/temp_measurements", routes.GetTempMeasurements)
	router.GET("/temp_measurements/:device_name/:timestamp", routes.GetTempMeasurement)
	router.POST("/temp_measurement/:device_name", routes.PostTempMeasurement)
	router.DELETE("/temp_measurement/:device_name/:timestamp", routes.DeleteTempMeasurement)

	router.Run(":8080")
}
