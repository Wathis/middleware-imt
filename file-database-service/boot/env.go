package boot

import (
	"log"
	"os"
)

var (
	BrokerUrl    = "localhost"
	BrokerPort   = "1883"
	CsvDataPath  = "./sensorsData/"
	MqttClientID = "sub2"
	MqttTopic    = "sensor/measure"
)

func LoadEnv() {
	LoadEnvVariable("BROKER_URL", &BrokerUrl)
	LoadEnvVariable("BROKER_PORT", &BrokerPort)
	LoadEnvVariable("CSV_EXPORT_PATH", &CsvDataPath)
	LoadEnvVariable("MQTT_CLIENT_ID", &MqttClientID)
	LoadEnvVariable("MQTT_TOPIC", &MqttTopic)
}

func LoadEnvVariable(key string, variable *string) {
	content, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Can't load variable %s, using default value %s", key, *variable)
		return
	}
	*variable = content
}
