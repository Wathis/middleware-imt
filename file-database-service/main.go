package main

import (
	"encoding/json"
	"file-database-service/csvwritter"
	"file-database-service/mqtttools"
	"file-database-service/sensor"
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const CsvDataPath = "./sensorsData/"
const MqttBrokerURI = "tcp://localhost:1883"
const MqttClientID = "sub2"
const MqttTopic = "sensor/mesure"

// const testData = `{"SensorID": 43, "AirportID": "NTE", "mesureType": "temp", "mesureValue": 17.33, "timestamp": 1570731034}`

func main() {
	client := mqtttools.Connect(MqttBrokerURI, MqttClientID)
	var wg sync.WaitGroup
	wg.Add(1)

	data := sensor.SensorData{}
	client.Subscribe(MqttTopic, 0, func(client mqtt.Client, msg mqtt.Message) {
		// Each time we receive data from broker
		json.Unmarshal([]byte(msg.Payload()), &data)
		fmt.Println("Data received: ", data)
		csvwritter.AddDataToCsv(data, CsvDataPath)
	})

	wg.Wait()
}
