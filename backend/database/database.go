package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	devices          *mongo.Collection
	tempMeasurements *mongo.Collection
	bubMeasurements  *mongo.Collection
)

// Device information
type Device struct {
	Name    string   `json:"name" bson:"_id"`
	Censors []string `json:"censors" bson:"censors"`
}

// TempMeasurement information
type TempMeasurement struct {
	Device          Device  `bson:"device" json:"device"`
	Timestamp       int64   `bson:"timestamp" json:"timestamp"`
	Value           float64 `json:"value" bson:"value"`
	MeasurementUnit string  `json:"measurement_unit" bson:"measurement_unit"`
}

// BubMeasurement information
type BubMeasurement struct {
	Device    Device `bson:"device" json:"device"`
	Timestamp int64  `bson:"timestamp" json:"timestamp"`
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

	devices = client.Database("iot").Collection("devices")
	tempMeasurements = client.Database("iot").Collection("tempMeasurements")
	bubMeasurements = client.Database("iot").Collection("bubMeasurements")
}

// GetAllDevices returns all devices from database
func GetAllDevices(c context.Context) ([]Device, error) {
	var returnDevices []Device
	cur, err := devices.Find(c, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(c)
	for cur.Next(c) {
		var device Device
		err := cur.Decode(&device)
		if err != nil {
			panic(err)
		}
		returnDevices = append(returnDevices, device)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}

	if returnDevices == nil {
		return []Device{}, errors.New("Devices not found")
	}

	return returnDevices, nil
}

// GetDevice returns one device from database
func GetDevice(c context.Context, name string) (Device, error) {
	var device Device

	err := devices.FindOne(c, bson.M{"_id": name}).Decode(&device)
	if err != nil {
		var notFoundErrorMsg = "mongo: no documents in result"
		if err.Error() == notFoundErrorMsg {
			return Device{}, errors.New("Device not found")
		}
		panic(err)
	}

	return device, nil
}

// InsertDevice adds one device to database
func InsertDevice(c context.Context, device Device) error {
	_, err := devices.InsertOne(c, device)
	if err != nil {
		merr := err.(mongo.WriteException)
		errCode := merr.WriteErrors[0].Code
		if errCode == 11000 {
			return errors.New("Device already exists")
		}
		panic(err)
	}

	return nil
}

// DeleteDevice returns one device from database
func DeleteDevice(c context.Context, name string) error {
	res, err := devices.DeleteOne(c, bson.M{"_id": name})
	if err != nil {
		panic(err)
	}
	if res.DeletedCount == 0 {
		return errors.New("Device not found")
	}

	return nil
}

// GetAllBubMeasurements returns all bubble measurements from database
func GetAllBubMeasurements(c context.Context) ([]BubMeasurement, error) {
	var returnMeasurements []BubMeasurement
	cur, err := bubMeasurements.Find(c, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(c)
	for cur.Next(c) {
		var bubMeasurement BubMeasurement
		err := cur.Decode(&bubMeasurement)
		if err != nil {
			panic(err)
		}
		returnMeasurements = append(returnMeasurements, bubMeasurement)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}

	if returnMeasurements == nil {
		return []BubMeasurement{}, errors.New("Measurements not found")
	}

	return returnMeasurements, nil
}

// GetBubMeasurements returns bubble measurements between timeframe from database
func GetBubMeasurements(c context.Context, deviceName string,
	startTime int64, endTime int64) ([]BubMeasurement, error) {
	var returnMeasurements []BubMeasurement

	device, err := GetDevice(c, deviceName)
	if err != nil {
		return []BubMeasurement{}, errors.New("Measurement not found")
	}

	filter := bson.M{"timestamp": bson.M{"$gt": startTime, "$lt": endTime},
		"device": device}
	cur, err := bubMeasurements.Find(c, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(c)
	for cur.Next(c) {
		var bubMeasurement BubMeasurement
		err := cur.Decode(&bubMeasurement)
		if err != nil {
			panic(err)
		}
		returnMeasurements = append(returnMeasurements, bubMeasurement)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}

	if returnMeasurements == nil {
		return []BubMeasurement{}, errors.New("Measurements not found")
	}

	return returnMeasurements, nil
}

// InsertBubMeasurement adds one bubble measurement to database
func InsertBubMeasurement(c context.Context, bubMeasurement BubMeasurement) error {
	_, err := bubMeasurements.InsertOne(c, bubMeasurement)
	if err != nil {
		panic(err)
	}

	return nil
}

// DeleteBubMeasurements deletes bubble measurements between timeframe from database
func DeleteBubMeasurements(c context.Context, deviceName string,
	startTime int64, endTime int64) error {
	var device Device

	device, err := GetDevice(c, deviceName)
	if err != nil {
		return errors.New("Device not found")
	}

	filter := bson.M{"timestamp": bson.M{"$gt": startTime, "$lt": endTime},
		"device": device}
	res, err := bubMeasurements.DeleteMany(c, filter)
	if err != nil {
		panic(err)
	}
	if res.DeletedCount == 0 {
		return errors.New("Device not found")
	}

	return nil
}

// GetAllTempMeasurements returns all temp measurements from database
func GetAllTempMeasurements(c context.Context) ([]TempMeasurement, error) {
	var returnMeasurements []TempMeasurement
	cur, err := tempMeasurements.Find(c, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(c)
	for cur.Next(c) {
		var tempMeasurement TempMeasurement
		err := cur.Decode(&tempMeasurement)
		if err != nil {
			panic(err)
		}
		returnMeasurements = append(returnMeasurements, tempMeasurement)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}

	if returnMeasurements == nil {
		return []TempMeasurement{}, errors.New("Measurements not found")
	}

	return returnMeasurements, nil
}

// GetTempMeasurements returns temp measurements between timeframe from database
func GetTempMeasurements(c context.Context, deviceName string,
	startTime int64, endTime int64) ([]TempMeasurement, error) {
	var returnMeasurements []TempMeasurement

	device, err := GetDevice(c, deviceName)
	if err != nil {
		return []TempMeasurement{}, errors.New("Measurement not found")
	}

	filter := bson.M{"timestamp": bson.M{"$gt": startTime, "$lt": endTime},
		"device": device}
	cur, err := tempMeasurements.Find(c, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(c)
	for cur.Next(c) {
		var tempMeasurement TempMeasurement
		err := cur.Decode(&tempMeasurement)
		if err != nil {
			panic(err)
		}
		returnMeasurements = append(returnMeasurements, tempMeasurement)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}

	if returnMeasurements == nil {
		return []TempMeasurement{}, errors.New("Measurements not found")
	}

	return returnMeasurements, nil
}

// InsertTempMeasurement adds one temp measurement to database
func InsertTempMeasurement(c context.Context, tempMeasurement TempMeasurement) error {
	_, err := tempMeasurements.InsertOne(c, tempMeasurement)
	if err != nil {
		panic(err)
	}

	return nil
}

// DeleteTempMeasurements deletes bubble measurements between timeframe from database
func DeleteTempMeasurements(c context.Context, deviceName string,
	startTime int64, endTime int64) error {
	var device Device

	device, err := GetDevice(c, deviceName)
	if err != nil {
		return errors.New("Device not found")
	}

	filter := bson.M{"timestamp": bson.M{"$gt": startTime, "$lt": endTime},
		"device": device}
	res, err := tempMeasurements.DeleteMany(c, filter)
	if err != nil {
		panic(err)
	}
	if res.DeletedCount == 0 {
		return errors.New("Device not found")
	}

	return nil
}
