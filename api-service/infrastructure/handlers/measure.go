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
	Averages map[string]float64 `json:"averages"`
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
	from, err := strconv.ParseInt(r.URL.Query().Get("from"), 10, 64)
	if err != nil {
		lib.RespondWithError(rw, err, http.StatusBadRequest)
		return
	}
	to, err := strconv.ParseInt(r.URL.Query().Get("to"), 10, 64)
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

// HandleMeasureAverages godoc
// @Description Measure average for day timestamp
// @Produce  json
// @Param day_timestamp path float64 true "Second timestamp of the day chosen (any timestamp during the day works)"
// @Success 200 {object} handlers.AverageResponse
// @Router /measures/{day_timestamp}/average [get]
func HandleMeasureAverages(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dayChosenTimestamp, err := strconv.ParseInt(vars["day_timestamp"], 10, 64)
	if err != nil {
		lib.RespondWithError(rw, err, http.StatusBadRequest)
		return
	}
	averages, err := application.MeasureRepository.FindMeasureAveragesForDay(dayChosenTimestamp)
	if err != nil {
		lib.RespondWithError(rw, err, http.StatusInternalServerError)
		return
	}
	lib.RespondWithJson(rw, AverageResponse{Averages: averages}, http.StatusOK)
}
