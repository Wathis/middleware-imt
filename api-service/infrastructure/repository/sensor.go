package repository

import (
	"api-service/application"
	"api-service/domain"
	"fmt"
	"github.com/pkg/errors"
)

type SensorRepository struct {
}

func NewSensorRepository() *SensorRepository {
	return &SensorRepository{}
}

func (s *SensorRepository) FindSensorMeasures(sensorId int) (measures []domain.Measure, err error) {
	keys, _, err := application.RedisClient.Scan(0, fmt.Sprintf("sensor:%d*", sensorId), 0).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Can't scan sensors")
	}
	measures, err = getMeasuresForKeys(keys)
	if err != nil {
		return nil, err
	}
	return measures, nil
}
