package boot

import (
	"fmt"
	"log"
	"redis-database-service/application"
	"redis-database-service/infrastructure/repository"

	"github.com/go-redis/redis/v7"
)

func LoadRepositories() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", RedisURI, RedisPort),
		DB:   0, // use default DB
	})
	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Panicf("Error ping redis : %s", err.Error())
	}
	if pong != "PONG" {
		log.Panicf("Error ping redis, response instead of pong : %s", pong)
	}

	// Initialisation des connecteurs vers la base et des interfaces (DAO)
	application.RedisClient = redisClient
	application.MeasureRepository = repository.NewMeasureRepository()
}
