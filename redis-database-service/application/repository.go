package application

import (
	"redis-database-service/domain"

	redis "github.com/go-redis/redis/v7"
)

var (
	RedisClient       *redis.Client
	MeasureRepository MeasureRepositoryInterface
)

type MeasureRepositoryInterface interface {
	SaveMeasure(measure domain.Measure) error
}
