package repository

import (
	"encoding/json"
	"fmt"
	"redis-database-service/cmd/application"
	entities "redis-database-service/internal/entities"
	"strconv"

	redis "github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

type MeasureRepository struct {
}

func NewMeasureRepository() *MeasureRepository {
	return &MeasureRepository{}
}

func (m *MeasureRepository) SaveMeasure() error {
	saveMeasureJSON(measure)
	saveMeasureTimestamp(measure)
	saveMeasureValue(measure)
}

func saveMeasureValue(measure domain.Measure) error {
	// TODO : Prendre en compte le nombre de ligne retourné et mettre une erreur en conséquence
	application.RedisClient.Zadd(getKeyZaddValue(measure), &redis.Z{
		Score:  float64(measure.MeasureValue),
		Member: getKeySet(measure),
	})
}

func saveMeasureTimestamp(measure domain.Measure) error {
	// TODO : Prendre en compte le nombre de ligne retourné et mettre une erreur en conséquence
	application.RedisClient.ZAdd(getKeyZaddTimestamp(measure), &redis.Z{
		Score:  float64(measure.Timestamp),
		Member: getKeySet(measure),
	})
}

func saveMeasureJSON(measure domain.Measure) error {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "error saveMeasureJSON:json.Marshal")
	}
	fmt.Println("Valeur : " + fmt.Sprintf("%s", dataJSON))
	application.RedisClient.Set(getKeySet(measure), fmt.Sprintf("%s", dataJSON), 0)
}

func getKeySet(measure entities.Measure) string {
	key := fmt.Sprintf("sensor:%s:measure:%s", strconv.Itoa(measure.SensorID), strconv.Itoa(measure.Timestamp))
	fmt.Println("Key : " + key)
	return key
}

func getKeyZaddValue(measure entities.Measure) string {
	key := fmt.Sprintf("measure_value:%s", measure.MeasureType)
	return key
}

func getKeyZaddTimestamp(measure entities.Measure) string {
	key := fmt.Sprintf("measure_timestamp:%s", measure.MeasureType)
	return key
}
