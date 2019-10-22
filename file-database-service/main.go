package main

import (
	"encoding/json"
	"file-database-service/boot"
	"file-database-service/csvwritter"
	"file-database-service/mqtttools"
	"file-database-service/sensor"
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// const testData = `{"SensorID": 43, "AirportID": "NTE", "mesureType": "temp", "mesureValue": 17.33, "timestamp": 1570731034}`

func main() {

	boot.LoadEnv()

	client := mqtttools.Connect("tcp://"+boot.BrokerUrl+":"+boot.BrokerPort, boot.MqttClientID)
	var wg sync.WaitGroup
	wg.Add(1)

	data := sensor.SensorData{}
	client.Subscribe(boot.MqttTopic, 0, func(client mqtt.Client, msg mqtt.Message) {
		// Each time we receive data from broker
		json.Unmarshal([]byte(msg.Payload()), &data)
		fmt.Println("Data received: ", data)
		csvwritter.AddDataToCsv(data, boot.CsvDataPath)
	})

	wg.Wait()
}
