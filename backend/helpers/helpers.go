package helpers

import "strconv"

// TempMeasurementRead is used for reading json from posting temp measurement
type TempMeasurementRead struct {
	Value           string `json:"value" bson:"value"`
	MeasurementUnit string `json:"measurement_unit" bson:"measurement_unit"`
}

// StringToIntConverter converts string to int64
func StringToIntConverter(timestamp string) int64 {
	n, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

// StringToFloatConverter converts string to float64
func StringToFloatConverter(stringToConvert string) float64 {
	n, err := strconv.ParseFloat(stringToConvert, 64)
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
