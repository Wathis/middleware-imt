package handlers

import (
	"api-service/application"
	"api-service/domain"
	"api-service/internal/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestHandleMeasures(t *testing.T) {
	measures := []domain.Measure{{
		SensorID:     1,
		AirportID:    "CDC",
		MeasureType:  "temperature",
		MeasureValue: 20,
		Timestamp:    1571229160,
	}}
	jsonMeasure, _ := json.Marshal(measures)
	var tests = []struct {
		name                  string
		measureRepositoryMock *mocks.MeasureRepositoryMock
		expectedCode          int
		expected              string
	}{
		{
			"Nominal case",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				measureRepositoryMock.On("FindMeasures").Return(measures, nil).Once()
				return measureRepositoryMock
			}(),
			http.StatusOK,
			string(jsonMeasure),
		},
		{
			"Error case",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				measureRepositoryMock.On("FindMeasures").Return(nil, errors.New("error while FindMeasures")).Once()
				return measureRepositoryMock
			}(),
			http.StatusInternalServerError,
			"error while FindMeasures",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application.MeasureRepository = tt.measureRepositoryMock
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			HandleMeasures(rec, req)
			if body := rec.Body.String(); !strings.Contains(body, tt.expected) {
				t.Errorf("Error body \nexpected contains: %s\ngot: %s", tt.expected, rec.Body.String())
			}
			if code := rec.Code; code != tt.expectedCode {
				t.Errorf("Error code \nexpected: %d\ngot: %d", tt.expectedCode, rec.Code)
			}
			tt.measureRepositoryMock.AssertExpectations(t)
		})
	}
}
func TestHandleMeasuresWithMeasureType(t *testing.T) {
	measures := []domain.Measure{{
		SensorID:     1,
		AirportID:    "CDC",
		MeasureType:  "temperature",
		MeasureValue: 20,
		Timestamp:    1571229160,
	}}
	jsonMeasure, _ := json.Marshal(measures)
	var tests = []struct {
		name                  string
		measureRepositoryMock *mocks.MeasureRepositoryMock
		param                 string
		vars                  map[string]string
		expectedCode          int
		expected              string
	}{
		{
			"Nominal case",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				measureRepositoryMock.On("FindMeasuresBetweenTimestamp", mock.Anything, mock.Anything, mock.Anything).Return(measures, nil).Once()
				return measureRepositoryMock
			}(),
			fmt.Sprintf("?from=%d&to=%d", 1571229160, 1571229165),
			map[string]string{
				"measure_type": "temperature",
			},
			http.StatusOK,
			string(jsonMeasure),
		},
		{
			"No params",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				return measureRepositoryMock
			}(),
			"",
			map[string]string{
				"measure_type": "temperature",
			},
			http.StatusBadRequest,
			"The from field is required",
		},
		{
			"No params",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				return measureRepositoryMock
			}(),
			"",
			map[string]string{
				"measure_type": "temperature",
			},
			http.StatusBadRequest,
			"The to field is required",
		},
		{
			"Wrong param format",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				return measureRepositoryMock
			}(),
			fmt.Sprintf("?from=%s&to=%s", "zdzd", "zd"),
			map[string]string{
				"measure_type": "temperature",
			},
			http.StatusBadRequest,
			"strconv.ParseInt: parsing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application.MeasureRepository = tt.measureRepositoryMock
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/test%s", tt.param), strings.NewReader(""))
			req = mux.SetURLVars(req, tt.vars)
			HandleMeasuresWithMeasureType(rec, req)
			if body := rec.Body.String(); !strings.Contains(body, tt.expected) {
				t.Errorf("Error body \nexpected contains: %s\ngot: %s", tt.expected, rec.Body.String())
			}
			if code := rec.Code; code != tt.expectedCode {
				t.Errorf("Error code \nexpected: %d\ngot: %d", tt.expectedCode, rec.Code)
			}

			tt.measureRepositoryMock.AssertExpectations(t)
		})
	}
}

func TestHandleMeasureAverage(t *testing.T) {
	measureType := "temperature"
	average := 10.2
	timestamp := time.Now().Unix()
	averages := map[string]float64{
		measureType: average,
	}
	AverageResponse := AverageResponse{Averages: averages}
	averageResponseJson, _ := json.Marshal(AverageResponse)
	var tests = []struct {
		name                  string
		measureRepositoryMock *mocks.MeasureRepositoryMock
		vars                  map[string]string
		expectedCode          int
		expected              string
	}{
		{
			"Nominal case",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				measureRepositoryMock.On("FindMeasureAveragesForDay", timestamp).Return(averages, nil).Once()
				return measureRepositoryMock
			}(),
			map[string]string{
				"day_timestamp": strconv.FormatInt(timestamp, 10),
			},
			http.StatusOK,
			string(averageResponseJson),
		},
		{
			"Error case",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				measureRepositoryMock.On("FindMeasureAveragesForDay", timestamp).Return(nil, errors.New("error while FindMeasureAveragesForDay")).Once()
				return measureRepositoryMock
			}(),
			map[string]string{
				"day_timestamp": strconv.FormatInt(timestamp, 10),
			},
			http.StatusInternalServerError,
			"error while FindMeasureAveragesForDay",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application.MeasureRepository = tt.measureRepositoryMock
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			req = mux.SetURLVars(req, tt.vars)
			HandleMeasureAverages(rec, req)
			if body := rec.Body.String(); !strings.Contains(body, tt.expected) {
				t.Errorf("Error body \nexpected contains: %s\ngot: %s", tt.expected, rec.Body.String())
			}
			if code := rec.Code; code != tt.expectedCode {
				t.Errorf("Error code \nexpected: %d\ngot: %d", tt.expectedCode, rec.Code)
			}
			tt.measureRepositoryMock.AssertExpectations(t)
		})
	}
}
