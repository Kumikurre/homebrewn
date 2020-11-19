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
	router.GET("/device/:id", routes.GetDevice)
	router.POST("/device", routes.PostDevice)

	router.Run(":8080")
}
