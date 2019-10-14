package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const file = "./config/config.json"

type Config struct {
	SensorTemp struct {
		BrokerURL  string `json:"brokerUrl"`
		BrokerPort int64  `json:"brokerPort"`
		SensorID   int64  `json:"sensorId"`
		SensorType string `json:"sensorType"`
		Qos        byte   `json:"qos"`
	} `json:"sensorTemp"`
	SensorWind struct {
		BrokerURL  string `json:"brokerUrl"`
		BrokerPort int64  `json:"brokerPort"`
		SensorID   int64  `json:"sensorId"`
		SensorType string `json:"sensorType"`
		Qos        byte   `json:"qos"`
	} `json:"sensorWind"`
	SensorPressure struct {
		BrokerURL  string `json:"brokerUrl"`
		BrokerPort int64  `json:"brokerPort"`
		SensorID   int64  `json:"sensorId"`
		SensorType string `json:"sensorType"`
		Qos        byte   `json:"qos"`
	} `json:"sensorPressure"`
}

func LoadConfiguration() Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
