package repository

import (
	"api-service/application"
	"api-service/domain"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type MeasureRepository struct {
}

func NewMeasureRepository() *MeasureRepository {
	return &MeasureRepository{}
}

func (m *MeasureRepository) FindMeasures() (measures []domain.Measure, err error) {
	var cursor uint64 = 0
	var nextCursor uint64 = 1
	var keys []string
	var keysForCurrentCursor []string
	for nextCursor != 0 {
		keysForCurrentCursor, nextCursor, err = application.RedisClient.Scan(cursor, "sensor:*", 0).Result()
		cursor = nextCursor
		keys = append(keys, keysForCurrentCursor...)
	}
	if err != nil {
		return nil, errors.Wrap(err, "can't scan sensors")
	}
	return getMeasuresForKeys(keys)
}

func (m *MeasureRepository) FindMeasuresBetweenTimestamp(measureType string, from int64, to int64) (measures []domain.Measure, err error) {
	return findMeasureBetweenTimestamp(measureType, from, to)
}
func (m *MeasureRepository) FindMeasureAveragesForDay(dayChosenTimestamp int64) (map[string]float64, error) {
	currentTime := time.Unix(dayChosenTimestamp, 0)
	dayStartTimestamp := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	dayEndTimestamp := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Unix()
	measureTypeKeys, _, err := application.RedisClient.Scan(0, "measure_timestamp:*", 0).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "can't scan measure_timestamp for timestamp : %i", dayChosenTimestamp)
	}
	measureAverages := make(map[string]float64)
	for _, measureTypeKey := range measureTypeKeys {
		measureType := strings.Split(measureTypeKey, ":")[1]
		measures, err := findMeasureBetweenTimestamp(measureType, dayStartTimestamp, dayEndTimestamp)
		if err != nil {
			return nil, err
		}
		if len(measures) > 0 {
			measureAverages[measureType] = getMeasuresAverage(measures)
		}
	}
	return measureAverages, nil
}

func getMeasuresForKeys(keys []string) (measures []domain.Measure, err error) {
	var measure domain.Measure
	for i := 0; i < len(keys); i++ {
		log.Print(keys[i])
		content, err := application.RedisClient.Get(keys[i]).Result()
		if err != nil {
			return nil, errors.Wrapf(err, "can't get content of %s", keys[i])
		}
		err = json.Unmarshal([]byte(content), &measure)
		if err != nil {
			return nil, errors.Wrapf(err, "can't unmarshal content of %s", keys[i])
		}
		measures = append(measures, measure)
	}
	return measures, nil
}
func getMeasuresAverage(measures []domain.Measure) float64 {
	total := 0.0
	if len(measures) == 0 {
		return 0
	}
	for _, measure := range measures {
		total += measure.MeasureValue
	}
	return total / float64(len(measures))
}

func findMeasureBetweenTimestamp(measureType string, from int64, to int64) ([]domain.Measure, error) {
	measureKey := fmt.Sprintf("measure_timestamp:%s", measureType)
	opts := &redis.ZRangeBy{
		Min:    strconv.FormatInt(from, 10),
		Max:    strconv.FormatInt(to, 10),
		Offset: 0,
		Count:  0,
	}
	keys, err := application.RedisClient.ZRangeByScore(measureKey, opts).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "can't get result of ZRangeByScore from key %s and options %v", measureKey, opts)
	}
	return getMeasuresForKeys(keys)
}
