package handlers

import (
	"api-service/infrastructure/repository"
	"api-service/lib"
	"log"
	"net/http"
)

func HandleMeasures(rw http.ResponseWriter, r *http.Request) {
	measures, err := repository.FindMeasures()
	if err != nil {
		log.Print(err)
		lib.RespondWithError(rw, err, http.StatusInternalServerError)
		return
	}
	lib.RespondWithJson(rw, measures, http.StatusOK)
}

func HandleSensorMeasures(rw http.ResponseWriter, r *http.Request) {

}
