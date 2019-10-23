package boot

import (
	"file-database-service/application"
	"file-database-service/infrastructure/handler"
	"fmt"
	"log"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(brokerUrl string, clientID string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerUrl)
	opts.SetClientID(clientID)
	return opts
}

func Connect() {
	fmt.Printf("Trying to connect (%s , %s) ...", brokerUrl, clientID)
	opts := createClientOptions(fmt.Sprintf("%s:%s", brokerUrl, brokerPort), clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	application.MqttClient = &client
}

// ListenBrocker : Function écoutant le brocker et insérant les données reçues, dans la base de données
func ListenBrocker() {
	var wg sync.WaitGroup
	Connect()
	wg.Add(1)
	fmt.Println("Waiting messages...")
	(*application.MqttClient).Subscribe(topicName, 0, handler.MeasureHandler)
	wg.Wait()
}
