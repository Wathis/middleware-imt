package repository

import (
	"encoding/json"
	"fmt"
	"redis-database-service/application"
	"redis-database-service/domain"
	"strconv"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

type MeasureRepository struct {
}

func NewMeasureRepository() *MeasureRepository {
	return &MeasureRepository{}
}

// SaveMeasure save all the data in the database
func (m *MeasureRepository) SaveMeasure(measure domain.Measure) error {
	err := saveMeasureJSON(measure)
	err = saveMeasureTimestamp(measure)
	err = saveMeasureValue(measure)
	return err
}

func saveMeasureValue(measure domain.Measure) error {
	err := application.RedisClient.ZAdd(fmt.Sprintf("measure_value:%s", measure.MeasureType), &redis.Z{
		Score:  float64(measure.MeasureValue),
		Member: getKeySet(measure),
	}).Err()
	if err != nil {
		return errors.Wrap(err, "error saveMeasureValue:ZAdd")
	}
	return nil
}

func saveMeasureTimestamp(measure domain.Measure) error {
	err := application.RedisClient.ZAdd(fmt.Sprintf("measure_timestamp:%s", measure.MeasureType), &redis.Z{
		Score:  float64(measure.Timestamp),
		Member: getKeySet(measure),
	}).Err()
	if err != nil {
		return errors.Wrap(err, "error saveMeasureTimestamp:ZAdd")
	}
	return nil
}

func saveMeasureJSON(measure domain.Measure) error {
	dataJSON, err := json.Marshal(measure)
	if err != nil {
		return errors.Wrap(err, "error saveMeasureJSON:json.Marshal")
	}
	application.RedisClient.Set(getKeySet(measure), fmt.Sprintf("%s", dataJSON), 0)
	return nil
}

func getKeySet(measure domain.Measure) string {
	key := fmt.Sprintf("sensor:%s:measure:%s", strconv.Itoa(measure.SensorID), strconv.Itoa(measure.Timestamp))
	return key
}
