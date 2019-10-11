package boot

import (
	"api-service/infrastructure/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func LoadRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/measures", handlers.HandleMeasures)
	r.HandleFunc("/sensors", handlers.HandleSensors)
	r.HandleFunc("/sensors/{sensor_id}", handlers.HandleSensor)
	r.HandleFunc("/sensors/{sensor_id}/measures", handlers.HandleSensorMeasures)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + serverPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Listening on port " + serverPort)
	log.Fatal(srv.ListenAndServe())
}
