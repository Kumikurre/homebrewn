package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	collection *mongo.Collection
)

// Device information
type Device struct {
	Name     string `json:"name" bson:"name"`
	DeviceID string `json:"device_id" bson:"device_id"`
}

// Init mongo database
func Init(uri string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	collection = client.Database("iot").Collection("devices")
}

// GetAllDevices returns all devices from database
func GetAllDevices(c context.Context) []Device {
	var devices []Device
	cur, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(c)
	for cur.Next(c) {
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

	return devices
}

// InsertDevice adds one device to database
func InsertDevice(c context.Context, device Device) *mongo.InsertOneResult {
	res, err := collection.InsertOne(c, device)
	if err != nil {
		panic(err)
	}

	return res
}
