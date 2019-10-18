package mocks

import (
	"api-service/domain"
	"github.com/stretchr/testify/mock"
)

type MeasureRepositoryMock struct {
	mock.Mock
}

func (m *MeasureRepositoryMock) FindMeasures() ([]domain.Measure, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Measure), nil
	} else {
		return nil, args.Error(1)
	}
}

func (m *MeasureRepositoryMock) FindMeasuresBetweenTimestamp(measureType string, from int64, to int64) ([]domain.Measure, error) {
	args := m.Called(measureType, from, to)
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Measure), nil
	} else {
		return nil, args.Error(1)
	}
}

func (m *MeasureRepositoryMock) FindMeasureAveragesForDay(timestamp int64) (map[string]float64, error) {
	args := m.Called(timestamp)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(map[string]float64), nil
	}
}

type SensorRepositoryMock struct {
	mock.Mock
}

func (m *SensorRepositoryMock) FindSensorMeasures(sensorId int) (measures []domain.Measure, err error) {
	args := m.Called(sensorId)
	if args.Get(0) != 0 {
		return args.Get(0).([]domain.Measure), nil
	} else {
		return nil, args.Error(1)
	}
}
