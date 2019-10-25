package application

import (
	"file-database-service/domain"
)

var (
	MeasureRepository MeasureRepositoryInterface
)

type MeasureRepositoryInterface interface {
	SaveMeasure(measure domain.Measure) error
}
