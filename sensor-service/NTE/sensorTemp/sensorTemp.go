package main

import (
	"fmt"
	"strconv"
	"time"

	common "common"
	config "config"
	mqtt "mqtt"
)

const topic = "sensor/measure"

func main() {
	//Récupération données du fichier de config
	config := config.LoadConfiguration("configSensorTemp.json")
	//Connexion client MQTT
	client := mqtt.Connect(config.BrokerURL+":"+strconv.FormatInt(config.BrokerPort, 10), strconv.FormatInt(config.SensorID, 10))

	min := -10.0
	max := 40.0
	variation := 1.0
	for x := range time.Tick(time.Duration(config.IntervalSendDataSensor) * time.Second) {
		fmt.Println(x)

		sensorData, value := common.RandomSensorData(config.AiportID, min, max, config)
		min, max = common.ChangeMinMaxValues(value, variation)
		client.Publish(topic, config.Qos, false, sensorData)
	}

	client.Disconnect(200)
}
