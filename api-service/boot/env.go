package boot

import (
	"log"
	"os"
)

var (
	redisUrl  = "localhost"
	redisPort = "6379"

	serverPort = "8080"
)

func LoadEnv() {
	LoadEnvVariable("REDIS_URL", &redisUrl)
	LoadEnvVariable("REDIS_PORT", &redisPort)
	LoadEnvVariable("SERVER_PORT", &serverPort)
}

func LoadEnvVariable(key string, variable *string) {
	content, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Can't load variable %s, using default value %s", key, *variable)
		return
	}
	*variable = content
}
