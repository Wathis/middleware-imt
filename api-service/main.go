package main

import (
	"api-service/boot"
	"log"
)

func main() {
	log.Print("starting api-service")
	log.Print("Loading env")
	boot.LoadEnv()
	log.Print("Loading repositories")
	boot.LoadRepositories()
	log.Print("Loading http router")
	boot.LoadRouter()

}
