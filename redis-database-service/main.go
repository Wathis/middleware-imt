package main

import (
	"log"
	"redis-database-service/cmd/application"
	"redis-database-service/cmd/boot"
)

func main() {
	log.Print("starting redis-database-service")
	log.Print("Loading env")
	boot.LoadEnv()
	log.Print("Starting Redis database")
	application.LaunchRedisServer()
	log.Print("Listening Brocker...")
	application.ListenBrocker()
}
