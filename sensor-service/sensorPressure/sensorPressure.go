package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"../common"
	"../config"
	"../mqtt"
)

const topic = "sensor/measure"

var airports = []struct {
	name string
	min  float64
	max  float64
}{
	{name: "NTE", min: 0.0, max: 30.0},
	{name: "LYS", min: 0.0, max: 30.0},
	{name: "CDG", min: 0.0, max: 30.0},
}

func main() {
	//Récupération données du fichier de config
	config := config.LoadConfiguration("configSensorPressure.json")
	//Connexion client MQTT
	client := mqtt.Connect(config.BrokerURL+":"+strconv.FormatInt(config.BrokerPort, 10), strconv.FormatInt(config.SensorID, 10))

	variation := 1.0
	for x := range time.Tick(time.Duration(config.IntervalSendDataSensor) * time.Second) {
		fmt.Println(x)

		//Pick random airport
		rand.Seed(time.Now().UTC().UnixNano())
		rAirport := rand.Intn(3)
		fmt.Println(airports[rAirport])

		sensorData, value := common.RandomSensorData(airports[rAirport].name, airports[rAirport].min, airports[rAirport].max, config)
		airports[rAirport].min, airports[rAirport].max = common.ChangeMinMaxValues(value, variation)
		client.Publish(topic, config.Qos, false, sensorData)
	}

	client.Disconnect(200)
}
