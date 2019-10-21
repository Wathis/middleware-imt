package common

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"../config"
)

func RandomSensorData(name string, min float64, max float64, config config.Config) (string, float64) {
	//Generate seed to get a new random number
	rand.Seed(time.Now().UTC().UnixNano())
	rValue := min + rand.Float64()*(max-min)

	//JSON to return
	sensorData := `{
		"SensorID":` + strconv.FormatInt(config.SensorID, 10) + `,
		"AirportID":` + name + `,
		"MesureType":"` + config.SensorType + `",
		"MesureValue":` + strconv.FormatFloat((math.Floor(rValue*100)/100), 'f', 6, 64) + `,
		"Timestamp":` + strconv.FormatInt(time.Now().Unix(), 10) + `
	}`

	return sensorData, rValue
}

func ChangeMinMaxValues(value float64, variation float64) (float64, float64) {
	min := value - variation
	max := value + variation
	return min, max
}