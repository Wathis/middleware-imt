package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"redis-database-service/cmd/dao"
	entities "redis-database-service/internal/entities"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// Save : Initialise la connexion et insère l'entité Mesure dans la base de données
func Save(data entities.Mesure) {
	conn := dao.Connexion.Get()
	defer conn.Close()
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	key := getKeySet(data)
	fmt.Println("Valeur : " + fmt.Sprintf("%s", dataJSON))
	doCommand(conn, "SET", key, fmt.Sprintf("%s", dataJSON))

	keySet := getKeyZaddTimestamp(data)
	fmt.Println("keySet : " + keySet + " Valeur : " + fmt.Sprintf("%d", data.Timestamp) + " key :" + key)
	doCommand(conn, "ZADD", keySet, data.Timestamp, key)
	keySet = getKeyZaddValue(data)
	fmt.Println("keySet : " + keySet + " Valeur : " + fmt.Sprintf("%f", data.MesureValue) + " key :" + key)
	doCommand(conn, "ZADD", keySet, int(data.MesureValue), key)

}

func doCommand(conn redis.Conn, command string, key string, args ...interface{}) {
	_, err := conn.Do(command, key, args)
	if err != nil {
		log.Fatal(err)
	}
}

func getKeySet(mesure entities.Mesure) string {
	key := "sensor"
	key += ":" + strconv.Itoa(mesure.SensorID)
	key += ":measure"
	key += ":" + strconv.Itoa(mesure.Timestamp)
	fmt.Println("Key : " + key)
	return key

}

func getKeyZaddValue(mesure entities.Mesure) string {
	key := "measure_value"
	key += ":" + mesure.MesureType
	return key
}

func getKeyZaddTimestamp(mesure entities.Mesure) string {
	key := "measure_timestamp"
	key += ":" + mesure.MesureType
	return key
}
