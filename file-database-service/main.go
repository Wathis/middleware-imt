package main

import (
	"file-database-service/boot"
	"log"
)

// const testData = `{"SensorID": 43, "AirportID": "NTE", "mesureType": "temp", "mesureValue": 17.33, "timestamp": 1570731034}`

func main() {

	log.Print("starting file-database-service")
	log.Print("Loading env")
	boot.LoadEnv()
	log.Print("Loading repository")
	boot.LoadRepositories()
	log.Print("Listening Brocker...")
	boot.ListenBrocker()
	boot.LoadEnv()

}
