package boot

import (
	"api-service/application"
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
)

func LoadRepositories() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisUrl, redisPort),
		DB:   0, // use default DB
	})
	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Panicf("Error ping redis : %s", err.Error())
	}
	if pong != "PONG" {
		log.Panicf("Error ping redis, response instead of pong : %s", pong)
	}
	application.RedisClient = redisClient
}
