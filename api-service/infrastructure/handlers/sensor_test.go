package handlers

import (
	"api-service/application"
	"api-service/domain"
	"api-service/internal/mocks"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestHandleSensorMeasures(t *testing.T) {
	sensorId := 1
	measures := []domain.Measure{{
		SensorID:     sensorId,
		AirportID:    "CDC",
		MeasureType:  "temperature",
		MeasureValue: 20,
		Timestamp:    1571229160,
	}}
	jsonMeasure, _ := json.Marshal(measures)
	var tests = []struct {
		name                 string
		sensorRepositoryMock *mocks.SensorRepositoryMock
		vars                 map[string]string
		expectedCode         int
		expected             string
	}{
		{
			"Nominal case",
			func() *mocks.SensorRepositoryMock {
				sensorRepositoryMock := new(mocks.SensorRepositoryMock)
				sensorRepositoryMock.On("FindSensorMeasures", sensorId).Return(measures, nil).Once()
				return sensorRepositoryMock
			}(),
			map[string]string{
				"sensor_id": strconv.Itoa(sensorId),
			},
			http.StatusOK,
			string(jsonMeasure),
		},
		{
			"Error type case",
			func() *mocks.SensorRepositoryMock {
				sensorRepositoryMock := new(mocks.SensorRepositoryMock)
				return sensorRepositoryMock
			}(),
			map[string]string{
				"sensor_id": "wrong integer",
			},
			http.StatusBadRequest,
			"can't parse sensor_id as integer",
		},
		{
			"Internal error case",
			func() *mocks.SensorRepositoryMock {
				sensorRepositoryMock := new(mocks.SensorRepositoryMock)
				sensorRepositoryMock.On("FindSensorMeasures", sensorId).Return(nil, errors.New("error during FindSensorMeasures")).Once()
				return sensorRepositoryMock
			}(),
			map[string]string{
				"sensor_id": strconv.Itoa(sensorId),
			},
			http.StatusInternalServerError,
			"error during FindSensorMeasures",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application.SensorRepository = tt.sensorRepositoryMock
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			req = mux.SetURLVars(req, tt.vars)
			HandleSensorMeasures(rec, req)
			if body := rec.Body.String(); !strings.Contains(body, tt.expected) {
				t.Errorf("Error body \nexpected contains: %s\ngot: %s", tt.expected, rec.Body.String())
			}
			if code := rec.Code; code != tt.expectedCode {
				t.Errorf("Error code \nexpected: %d\ngot: %d", tt.expectedCode, rec.Code)
			}
			tt.sensorRepositoryMock.AssertExpectations(t)
		})
	}
}
