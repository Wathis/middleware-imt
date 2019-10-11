package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Error interface{} `json:"error"`
}

func RespondWithJson(rw http.ResponseWriter, payload interface{}, httpStatus int)  {
	body,err := json.Marshal(payload)
	if err != nil {
		log.Print(err)
		return
	}
	rw.WriteHeader(httpStatus)
	_,_ = rw.Write(body)
}

func RespondWithError(rw http.ResponseWriter, error interface{}, httpStatus int)  {
	body,err := json.Marshal(Error{Error:error})
	if err != nil {
		log.Print(err)
		return
	}
	rw.WriteHeader(httpStatus)
	_,_ = rw.Write(body)
}