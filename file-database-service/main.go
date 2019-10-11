package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const CsvDataPath = "./sensorsData/"
const MqttBrokerURI = "tcp://localhost:1883"
const MqttClientID = "sub2"
const MqttTopic = "sensor/mesure"

// const testData = `{"SensorID": 43, "AirportID": "NTE", "mesureType": "temp", "mesureValue": 17.33, "timestamp": 1570731034}`

func main() {
	client := connect(MqttBrokerURI, MqttClientID)
	var wg sync.WaitGroup
	wg.Add(1)

	data := sensorData{}
	client.Subscribe(MqttTopic, 0, func(client mqtt.Client, msg mqtt.Message) {
		// Each time we receive data from broker
		json.Unmarshal([]byte(msg.Payload()), &data)
		fmt.Println(data)
		addDataToCsv(data)
	})

	wg.Wait()
}

func addDataToCsv(row sensorData) {
	rowDatetime := getTime(row.Timestamp)

	var filename string = string(row.AirportID) + "-" + rowDatetime[:10] + "-" + string(row.MesureType) + ".csv"

	f, err := os.OpenFile(CsvDataPath+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	// Write line here
	writer.Write([]string{
		strconv.FormatInt(row.SensorID, 10),
		row.AirportID,
		row.MesureType,
		strconv.FormatFloat(row.MesureValue, 'f', 6, 32),
		rowDatetime,
	})

	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
}

func getTime(input int64) string {
	var tm time.Time = time.Unix(input, 0)
	var res string = tm.Format("2006-01-02 15:04:05")
	return res
}

func createClientOptions(brokerURI string, clientID string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientID)
	return opts
}

func connect(brokerURI string, clientID string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientID + ")...")
	opts := createClientOptions(brokerURI, clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

type sensorData struct {
	SensorID    int64
	AirportID   string
	MesureType  string
	MesureValue float64
	Timestamp   int64
}
