package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	IntervalSendDataSensor int64   `json:"intervalSendDataSensor"`
	BrokerURL              string  `json:"brokerUrl"`
	BrokerPort             int64   `json:"brokerPort"`
	AiportID               string  `json:"airportId"`
	SensorID               int64   `json:"sensorId"`
	SensorType             string  `json:"sensorType"`
	Qos                    byte    `json:"qos"`
	Min                    float64 `json:"min"`
	Max                    float64 `json:"max"`
	Variation              float64 `json:"variation"`
}

func LoadConfiguration(filename string) Config {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
