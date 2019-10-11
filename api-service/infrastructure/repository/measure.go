package repository

import (
	"api-service/application"
	"api-service/domain"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
)

func FindMeasures() (measures []domain.Measure, err error) {
	keys, _, err := application.RedisClient.Scan(0, "*", 1).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Can't scan sensors")
	}
	log.Print(keys)
	var measure domain.Measure

	for i := 0; i < len(keys); i++ {
		content, err := application.RedisClient.Get(keys[i]).Result()
		if err != nil {
			return nil, errors.Wrapf(err, "Can't get content of %s", keys[i])
		}
		err = json.Unmarshal([]byte(content), &measure)
		if err != nil {
			return nil, errors.Wrapf(err, "Can't unmarshal content of %s", keys[i])
		}
		measures = append(measures, measure)
	}
	return measures, nil
}
