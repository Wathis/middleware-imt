package boot

import (
	"redis-database-service/infrastructure/handler"
	"fmt"
	"log"
	"redis-database-service/application"
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
	fmt.Println("Waiting messages...")
	MqttClient.Subscribe(boot.TopicName, 0, handler.MeasureHandler)
	wg.Wait()
}