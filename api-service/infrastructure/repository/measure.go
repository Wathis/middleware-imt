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
)

type MeasureRepository struct {
}

func NewMeasureRepository() *MeasureRepository {
	return &MeasureRepository{}
}

func (m *MeasureRepository) FindMeasures() (measures []domain.Measure, err error) {
	keys, _, err := application.RedisClient.Scan(0, "sensor:*", 0).Result()
	if err != nil {
		return nil, errors.Wrap(err, "can't scan sensors")
	}
	return getMeasuresForKeys(keys)
}

func (m *MeasureRepository) FindMeasuresBetweenTimestamp(measureType string, from int, to int) (measures []domain.Measure, err error) {
	measureKey := fmt.Sprintf("measure_timestamp:%s", measureType)
	opts := &redis.ZRangeBy{
		Min:    strconv.Itoa(from),
		Max:    strconv.Itoa(to),
		Offset: 0,
		Count:  0,
	}
	keys, err := application.RedisClient.ZRangeByScore(measureKey, opts).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "can't get result of ZRangeByScore from key %s and options %v", measureKey, opts)
	}
	return getMeasuresForKeys(keys)
}
func (m *MeasureRepository) FindMeasureAverage(measureType string) (float64, error) {
	measureKey := fmt.Sprintf("measure_value:%s", measureType)
	keysWithScore, err := application.RedisClient.ZRangeWithScores(fmt.Sprintf("measure_value:%s", measureType), 0, -1).Result()
	if err != nil {
		return -1, errors.Wrapf(err, "can't get result of ZRange for key : %v", measureKey)
	}
	if len(keysWithScore) == 0 {
		return 0, nil
	}
	total := 0.0
	for i := 0; i < len(keysWithScore); i++ {
		total += keysWithScore[i].Score
	}
	return total / float64(len(keysWithScore)), nil
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
