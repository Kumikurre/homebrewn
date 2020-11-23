package main

import (
	gin "github.com/gin-gonic/gin"

	"github.com/Kumikurre/homebrewn/backend/database"
	"github.com/Kumikurre/homebrewn/backend/routes"
)

func main() {
	database.Init("mongodb://mongo:27017")
	router := gin.Default()

	router.POST("/device", routes.PostDevice)
	router.GET("/devices", routes.GetDevices)
	router.GET("/device/:name", routes.GetDevice)
	router.DELETE("/device/:name", routes.DeleteDevice)

	router.POST("/bub_measurement/:device_name", routes.PostBubMeasurement)

	router.GET("/bub_measurements_all", routes.GetAllBubMeasurements)
	router.GET("/bub_measurements_all/from/:start_time", routes.GetAllBubMeasurements)
	router.GET("/bub_measurements_all/from/:start_time/to/:end_time", routes.GetAllBubMeasurements)

	router.GET("/bub_measurements/:device_name", routes.GetBubMeasurements)
	router.GET("/bub_measurements/:device_name/from/:start_time", routes.GetBubMeasurements)
	router.GET("/bub_measurements/:device_name/from/:start_time/to/:end_time", routes.GetBubMeasurements)

	router.DELETE("/bub_measurements/:device_name", routes.DeleteBubMeasurements)
	router.DELETE("/bub_measurements/:device_name/from/:start_time", routes.DeleteBubMeasurements)
	router.DELETE("/bub_measurements/:device_name/from/:start_time/to/:end_time", routes.DeleteBubMeasurements)

	router.POST("/temp_measurement/:device_name", routes.PostTempMeasurement)

	router.GET("/temp_measurements_all", routes.GetAllTempMeasurements)
	router.GET("/temp_measurements_all/from/:start_time", routes.GetAllTempMeasurements)
	router.GET("/temp_measurements_all/from/:start_time/to/:end_time", routes.GetAllTempMeasurements)

	router.GET("/temp_measurements/:device_name", routes.GetTempMeasurements)
	router.GET("/temp_measurements/:device_name/from/:start_time", routes.GetTempMeasurements)
	router.GET("/temp_measurements/:device_name/from/:start_time/to/:end_time", routes.GetTempMeasurements)

	router.DELETE("/temp_measurements/:device_name", routes.DeleteTempMeasurements)
	router.DELETE("/temp_measurements/:device_name/from/:start_time", routes.DeleteTempMeasurements)
	router.DELETE("/temp_measurements/:device_name/from/:start_time/to/:end_time", routes.DeleteTempMeasurements)

	router.Run(":8080")
}
