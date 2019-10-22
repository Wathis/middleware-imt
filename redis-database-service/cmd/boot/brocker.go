package boot

import (
	"fmt"
	"log"
	"redis-database-service/cmd/application"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(brokerURI string, clientID string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientID)
	return opts
}

func Connect() {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientID + ")...")
	opts := createClientOptions(brokerURI, clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	application.MqttClient := &client
}

// ListenBrocker : Function écoutant le brocker et insérant les données reçues, dans la base de données
func ListenBrocker() {
	
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("En attente de message...")
	data := entities.Measure{}
	MqttClient.Subscribe(boot.TopicName, 0, func(client mqtt.Client, msg mqtt.Message) {
		// Parse le JSON dans un objet à chaque reception d'un message sur le topic
		log.Println("Message reçu : " + string(msg.Payload()))
		json.Unmarshal([]byte(msg.Payload()), &data)
		// Ajoute la measure dans la base de données
		go redis.Save(data)
	})
	wg.Wait()
}