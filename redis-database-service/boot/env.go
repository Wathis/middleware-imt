package boot

import (
	"log"
	"os"
)

var (
	brokerUrl  = "tcp://localhost"
	brokerPort = "1883"
	redisURI   = "127.0.0.1"
	redisPort  = "6379"
	clientID   = "sub1"
	topicName  = "sensor/measure"
)

func LoadEnv() {
	LoadEnvVariable("BROKER_URL", &brokerUrl)
	LoadEnvVariable("BROKER_PORT", &brokerPort)
	LoadEnvVariable("REDIS_URL", &redisURI)
	LoadEnvVariable("REDIS_PORT", &redisPort)
	LoadEnvVariable("CLIENT_ID", &clientID)
	LoadEnvVariable("TOPIC_NAME", &topicName)
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
