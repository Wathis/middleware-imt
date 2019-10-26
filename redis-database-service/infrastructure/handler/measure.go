package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"redis-database-service/application"
	"redis-database-service/domain"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MeasureHandler handle the reception of the brocker's messages
func MeasureHandler(client mqtt.Client, msg mqtt.Message) {
	// Parse le JSON dans un objet à chaque reception d'un message sur le topic
	data := domain.Measure{}
	log.Println("Message reçu : " + string(msg.Payload()))
	err := json.Unmarshal([]byte(msg.Payload()), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Ajoute la measure dans la base de données
	go application.MeasureRepository.SaveMeasure(data)
}
