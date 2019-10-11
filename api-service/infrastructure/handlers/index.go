package handlers

import (
	"api-service/lib"
	"net/http"
)

type IndexResponse struct {
	Status string `json:"status"`
}

func Index(rw http.ResponseWriter,r *http.Request)  {
	lib.RespondWithJson(rw, IndexResponse{
		Status: "OK",
	}, http.StatusOK)
}