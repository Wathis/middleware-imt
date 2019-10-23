package handlers

import (
	"api-service/application"
	"api-service/lib"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// HandleSensorMeasures godoc
// @Description Retrieve sensor measures for sensor id
// @Produce  json
// @Param sensor_id path int true "Ex: 1"
// @Success 200 {object} domain.Measure
// @Router /sensors/{sensor_id}/measures [get]
func HandleSensorMeasures(rw http.ResponseWriter, r *http.Request) {
	sensorId, _ := strconv.Atoi(mux.Vars(r)["sensor_id"])
	measures, err := application.SensorRepository.FindSensorMeasures(sensorId)
	if err != nil {
		lib.RespondWithError(rw, err, http.StatusInternalServerError)
		return
	}
	lib.RespondWithJson(rw, measures, http.StatusOK)

}
