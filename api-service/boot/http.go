package boot

import (
	"api-service/infrastructure/handlers"
	"api-service/infrastructure/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func LoadRouter() {
	r := mux.NewRouter()

	r.Use(middleware.HttpLogger)

	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/measures", handlers.HandleMeasures)
	r.HandleFunc("/measures/{measure_type:[a-zA-Z]+}", handlers.HandleMeasuresWithMeasureType)
	r.HandleFunc("/measures/{measure_type:[a-zA-Z]+}/average", handlers.HandleMeasureAverage)
	r.HandleFunc("/sensors/{sensor_id:[0-9]+}/measures", handlers.HandleSensorMeasures)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + serverPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Listening on port " + serverPort)
	log.Fatal(srv.ListenAndServe())
}
