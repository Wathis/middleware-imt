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

func (m *MeasureRepositoryMock) FindMeasuresBetweenTimestamp(measureType string, from int, to int) ([]domain.Measure, error) {
	args := m.Called(measureType, from, to)
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Measure), nil
	} else {
		return nil, args.Error(1)
	}
}

func (m *MeasureRepositoryMock) FindMeasureAverage(measureType string) (float64, error) {
	args := m.Called(measureType)
	if args.Get(0) != 0 {
		return args.Get(0).(float64), nil
	} else {
		return -1, args.Error(1)
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
