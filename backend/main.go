package main

import (
	"context"
	"log"
	"net/http"
	"time"

	gin "github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Device information
type Device struct {
	Name     string `json:"name" bson:"name"`
	DeviceID string `json:"device_id" bson:"device_id"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	collection := client.Database("iot").Collection("devices")

	router := gin.Default()

	// List all devices
	router.GET("/devices", func(c *gin.Context) {
		var devices []Device
		cur, err := collection.Find(context.TODO(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(context.TODO())
		for cur.Next(context.TODO()) {
			var device Device
			err := cur.Decode(&device)
			if err != nil {
				log.Fatal(err)
			}
			devices = append(devices, device)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, devices)
	})

	// Get more specific data about one device
	router.GET("/devices/:name", func(c *gin.Context) {
		c.String(http.StatusOK, "get device data")
	})

	router.POST("/device", func(c *gin.Context) {
		var device Device
		c.BindJSON(&device)

		res, err := collection.InsertOne(context.TODO(), device)
		if err != nil {
			panic(err)
		}
		c.String(http.StatusOK, "device %v posted", res.InsertedID)
	})

	router.Run(":8080")
}
