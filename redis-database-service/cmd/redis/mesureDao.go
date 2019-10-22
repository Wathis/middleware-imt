package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"redis-database-service/cmd/application"
	entities "redis-database-service/internal/entities"

	redis "github.com/go-redis/redis/v7"
)

// mosquitto_pub -h 127.0.0.1 -p 1883 -t "sensor/measure" -m "{\"sensorId\":1,\"airportId\":\"CDG\",\"measureType\":\"temperature\",\"measureValue\":10.7,\"timestamp\":1570966444}"
// Save : Initialise la connexion et insère l'entité Measure dans la base de données
func Save(data entities.Measure) {
	conn := application.RedisClient
	defer conn.Close()
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	key := getKeySet(data)
	fmt.Println("Valeur : " + fmt.Sprintf("%s", dataJSON))
	conn.Set(key, fmt.Sprintf("%s", dataJSON), 0)

	keySet := getKeyZaddTimestamp(data)
	fmt.Println("keySet : " + keySet + " Valeur : " + fmt.Sprintf("%d", data.Timestamp) + " key :" + key)
	conn.ZAdd(keySet, &redis.Z{
		Score:  float64(data.Timestamp),
		Member: key,
	})

	keySet = getKeyZaddValue(data)
	fmt.Println("keySet : " + keySet + " Valeur : " + fmt.Sprintf("%f", data.MeasureValue) + " key :" + key)
	conn.ZAdd(keySet, &redis.Z{
		Score:  data.MeasureValue,
		Member: key,
	})
}
