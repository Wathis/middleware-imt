package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"../config"
	"../mqtt"
)

func main() {
	//Récupération données du fichier de config
	config := config.LoadConfiguration()
	//Connexion client MQTT
	client := mqtt.Connect(config.SensorWind.BrokerURL+":"+strconv.FormatInt(config.SensorWind.BrokerPort, 10), strconv.FormatInt(config.SensorWind.SensorID, 10))

	//Loop every 10s
	//Temp variables
	minTemp := -10.0
	maxTemp := 40.0
	variationTemp := 2.0
	//Temp variables
	minWind := 0.0
	maxWind := 200.0
	variationWind := 10.0
	//Temp variables
	minPressure := 0.0
	maxPressure := 30.0
	variationPressure := 1.0
	for x := range time.Tick(10 * time.Second) {
		fmt.Println(x)
		//Temp
		sTemp, rValueTemp := randomSensorTempData(minTemp, maxTemp, "SensorTemp", config)
		minTemp = rValueTemp - variationTemp
		maxTemp = rValueTemp + variationTemp
		client.Publish("sensor/measure", config.SensorTemp.Qos, false, sTemp)
		//Wind
		sWind, rValueWind := randomSensorTempData(minWind, maxWind, "SensorWind", config)
		minWind = rValueWind - variationWind
		maxWind = rValueWind + variationWind
		client.Publish("sensor/measure", config.SensorWind.Qos, false, sWind)
		//Pressure
		sPressure, rValuePressure := randomSensorTempData(minPressure, maxPressure, "SensorPressure", config)
		minPressure = rValuePressure - variationPressure
		maxPressure = rValuePressure + variationPressure
		client.Publish("sensor/measure", config.SensorPressure.Qos, false, sPressure)
	}

	client.Disconnect(200)
}

func randomSensorTempData(min float64, max float64, sensorType string, config config.Config) (string, float64) {
	//Generate seed to get a new random number
	rand.Seed(time.Now().UTC().UnixNano())
	rValue := min + rand.Float64()*(max-min)

	//Access dynamically the fields
	r := reflect.ValueOf(config)
	fields := reflect.Indirect(r).FieldByName(sensorType)

	//JSON to return
	sensorData := `{
		"SensorID":` + strconv.FormatInt(fields.FieldByName("SensorID").Int(), 10) + `,
		"AirportID":"NTE",
		"MesureType":"` + fields.FieldByName("SensorType").String() + `",
		"MesureValue":` + strconv.FormatFloat((math.Floor(rValue*100)/100), 'f', 6, 64) + `,
		"Timestamp":` + strconv.FormatInt(time.Now().Unix(), 10) + `
	}`

	return sensorData, rValue
}
