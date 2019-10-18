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
	FindMeasuresBetweenTimestamp(string, int64, int64) ([]domain.Measure, error)
	FindMeasureAveragesForDay(int64) (map[string]float64, error)
}

type SensorRepositoryInterface interface {
	FindSensorMeasures(sensorId int) (measures []domain.Measure, err error)
}
