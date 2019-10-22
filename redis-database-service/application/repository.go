package application

import (
	"redis-database-service/domain"

	"github.com/go-redis/redis"
)

var (
	RedisClient       *redis.Client
	MeasureRepository MeasureRepositoryInterface
)

type MeasureRepositoryInterface interface {
	SaveMeasure(measure domain.Measure) error
}
