package application

import (
	"api-service/domain"
	"github.com/go-redis/redis/v7"
)

var (
	RedisClient       *redis.Client
	MeasureRepository MeasureRepositoryInterface
	SensorRepository  SensorRepositoryInterface
)

type MeasureRepositoryInterface interface {
	FindMeasures() ([]domain.Measure, error)
	FindMeasuresBetweenTimestamp(string, int, int) ([]domain.Measure, error)
	FindMeasureAverage(string) (float64, error)
}

type SensorRepositoryInterface interface {
	FindSensorMeasures(sensorId int) (measures []domain.Measure, err error)
}
