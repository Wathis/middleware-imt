package boot

import (
	"log"
	"os"
)

var (
	BrockerURI  = "tcp://localhost"
	BrockerPort = "1883"
	RedisURI    = "127.0.0.1"
	RedisPort   = "6379"
	ClientID    = "sub1"
	TopicName   = "sensor/measure"
)

func LoadEnv() {
	LoadEnvVariable("BROCKER_URL", &BrockerURI)
	LoadEnvVariable("BROCKER_PORT", &BrockerPort)
	LoadEnvVariable("REDIS_URL", &RedisURI)
	LoadEnvVariable("REDIS_PORT", &RedisPort)
	LoadEnvVariable("CLIENT_ID", &ClientID)
	LoadEnvVariable("TOPIC_NAME", &TopicName)
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
