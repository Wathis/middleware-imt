package application

import (
	"github.com/go-redis/redis"
)

var (
	RedisClient       *redis.Client
	MeasureRepository MeasureRepositoryInterface
)

type MeasureRepositoryInterface interface {
	SaveMeasure() error
}
