package application

import (
	"encoding/json"
	"fmt"
	"log"
	"redis-database-service/cmd/boot"
	connection "redis-database-service/cmd/brocker"
	"redis-database-service/cmd/dao"
	"redis-database-service/cmd/redis"
	entities "redis-database-service/internal/entities"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// ListenBrocker : Function écoutant le brocker et insérant les données reçues, dans la base de données
func ListenBrocker() {

	client := connection.Connect(boot.BrockerURI+":"+boot.BrockerPort, boot.ClientID)
	var wg sync.WaitGroup
	wg.Add(1)
	dao.InitDatabase(boot.RedisURI + ":" + boot.RedisPort)
	fmt.Println("En attente de message...")
	data := entities.Mesure{}
	client.Subscribe(boot.TopicName, 0, func(client mqtt.Client, msg mqtt.Message) {
		// Parse le JSON dans un objet à chaque reception d'un message sur le topic
		log.Println("Message reçu : " + string(msg.Payload()))
		json.Unmarshal([]byte(msg.Payload()), &data)
		// Ajoute la mesure dans la base de données
		go redis.Save(data)
	})
	wg.Wait()
}
