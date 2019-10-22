package boot

import (
	"file-database-service/internal/exportpath"
	"log"
	"os"
)

var (
	brokerUrl      = "tcp://localhost"
	brokerPort     = "1883"
	clientID       = "sub2"
	topicName      = "sensor/measure"
	fileExportPath = "C:/Users/cedri/Desktop/sensorMeasures"
)

func LoadEnv() {
	LoadEnvVariable("BROKER_URL", &brokerUrl)
	LoadEnvVariable("BROKER_PORT", &brokerPort)
	LoadEnvVariable("FILE_EXPORT_PATH", &fileExportPath)
	LoadEnvVariable("MQTT_CLIENT_ID", &clientID)
	LoadEnvVariable("MQTT_TOPIC", &topicName)

	exportpath.FileExportPath = fileExportPath
}

func LoadEnvVariable(key string, variable *string) {
	content, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Can't load variable %s, using default value %s", key, *variable)
		os.Setenv(key, *variable)
		return
	}
	*variable = content
}
