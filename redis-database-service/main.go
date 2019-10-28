package main

import (
	"log"
	"redis-database-service/boot"
)

func main() {
	log.Print("starting redis-database-service")
	log.Print("Loading env")
	boot.LoadEnv()
	log.Print("Loading repository")
	boot.LoadRepositories()
	log.Print("Listening Brocker...")
	boot.ListenBrocker()
}
