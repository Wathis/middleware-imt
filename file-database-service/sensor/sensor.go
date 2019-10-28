package sensor

type SensorData struct {
	SensorID     int64   `json:"sensorId"`
	AirportID    string  `json:"airportId"`
	MeasureType  string  `json:"measureType"`
	MeasureValue float64 `json:"measureValue"`
	Timestamp    int64   `json:"timestamp"`
}
