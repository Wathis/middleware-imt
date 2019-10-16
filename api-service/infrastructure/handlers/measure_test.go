package handlers

import (
	"api-service/application"
	"api-service/domain"
	"api-service/internal/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
				measureRepositoryMock.On("FindMeasures").Return(measures, nil)
				return measureRepositoryMock
			}(),
			http.StatusOK,
			string(jsonMeasure),
		},
		{
			"Error case",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				measureRepositoryMock.On("FindMeasures").Return(nil, errors.New("error while FindMeasures"))
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
				measureRepositoryMock.On("FindMeasuresBetweenTimestamp").Return(measures, nil)
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
			"Error case",
			func() *mocks.MeasureRepositoryMock {
				measureRepositoryMock := new(mocks.MeasureRepositoryMock)
				measureRepositoryMock.On("FindMeasuresBetweenTimestamp").Return(nil, errors.New("error while FindMeasures"))
				return measureRepositoryMock
			}(),
			fmt.Sprintf("from=%d?to=%d", 1571229160, 1571229165),
			map[string]string{},
			http.StatusBadRequest,
			"error while FindMeasures",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application.MeasureRepository = tt.measureRepositoryMock
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/test%s", tt.param), strings.NewReader(""))
			mux.SetURLVars(req, tt.vars)
			HandleMeasuresWithMeasureType(rec, req)
			if body := rec.Body.String(); !strings.Contains(body, tt.expected) {
				t.Errorf("Error body \nexpected contains: %s\ngot: %s", tt.expected, rec.Body.String())
			}
			if code := rec.Code; code != tt.expectedCode {
				t.Errorf("Error code \nexpected: %d\ngot: %d", tt.expectedCode, rec.Code)
			}
		})
	}
}

func TestHandleMeasureAverage(t *testing.T) {

}
