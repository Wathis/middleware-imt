package handlers

import (
	"api-service/application"
	"api-service/lib"
	"github.com/gorilla/mux"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"strconv"
)

type AverageResponse struct {
	Average float64 `json:"average"`
}

// HandleMeasures godoc
// @Description List all measures
// @Produce  json
// @Success 200 {array} domain.Measure
// @Router /measures [get]
func HandleMeasures(rw http.ResponseWriter, r *http.Request) {
	measures, err := application.MeasureRepository.FindMeasures()
	if err != nil {
		lib.RespondWithError(rw, err, http.StatusInternalServerError)
		return
	}
	lib.RespondWithJson(rw, measures, http.StatusOK)
}

// HandleMeasuresWithMeasureType godoc
// @Description List measures between timestamp
// @Produce  json
// @Param from query int true "timestamp"
// @Param to query int true "timestamp"
// @Param measure_type path string true "Ex: temperature"
// @Success 200 {array} domain.Measure
// @Router /measures/{measure_type} [get]
func HandleMeasuresWithMeasureType(rw http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"from": []string{"required"},
		"to":   []string{"required"},
	}
	if validationErrors := lib.ValidateQueryString(r, rules, nil); validationErrors != nil {
		lib.RespondWithJson(rw, validationErrors, http.StatusBadRequest)
		return
	}
	from, err := strconv.Atoi(r.URL.Query().Get("from"))
	if err != nil {
		lib.RespondWithError(rw, err, http.StatusBadRequest)
		return
	}
	to, err := strconv.Atoi(r.URL.Query().Get("to"))
	if err != nil {
		lib.RespondWithError(rw, err, http.StatusBadRequest)
		return
	}
	measures, err := application.MeasureRepository.FindMeasuresBetweenTimestamp(mux.Vars(r)["measure_type"], from, to)
	if err != nil {
		lib.RespondWithError(rw, err, http.StatusInternalServerError)
		return
	}
	lib.RespondWithJson(rw, measures, http.StatusOK)
}

// HandleMeasureAverage godoc
// @Description Measure average for measure type
// @Produce  json
// @Param measure_type path string true "Ex: temperature"
// @Success 200 {object} handlers.AverageResponse
// @Router /measures/{measure_type}/average [get]
func HandleMeasureAverage(rw http.ResponseWriter, r *http.Request) {
	average, err := application.MeasureRepository.FindMeasureAverage(mux.Vars(r)["measure_type"])
	if err != nil {
		lib.RespondWithError(rw, err, http.StatusInternalServerError)
		return
	}
	lib.RespondWithJson(rw, AverageResponse{Average: average}, http.StatusOK)
}
