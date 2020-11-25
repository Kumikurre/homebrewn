package helpers

import (
	"strconv"
	"time"
)

// TempMeasurementRead is used for reading json from posting temp measurement
type TempMeasurementRead struct {
	Value           *float64 `json:"value" bson:"value" binding:"required"`
	MeasurementUnit string   `json:"measurement_unit" bson:"measurement_unit" binding:"required"`
}

// StringToIntConverter converts string to int64
func StringToIntConverter(timestamp string) int64 {
	n, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

// Contains tells if string is in a slice
func Contains(strList []string, str string) bool {
	for _, v := range strList {
		if v == str {
			return true
		}
	}
	return false
}

// ParamReader reads timestamp imputs and converts them to int64
func ParamReader(startTimeInput string, endTimeInput string) (int64, int64) {
	var startTime int64
	var endTime int64
	if startTimeInput == "" {
		startTime = 0
	} else {
		startTime = StringToIntConverter(startTimeInput)
	}
	if endTimeInput == "" {
		endTime = time.Now().UnixNano()
	} else {
		endTime = StringToIntConverter(endTimeInput)
	}

	return startTime, endTime
}
