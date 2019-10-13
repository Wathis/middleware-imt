package domain

type Measure struct {
	SensorID     int     `json:"sensorId"`
	AirportID    string  `json:"airportId"`
	MeasureType  string  `json:"measureType"`
	MeasureValue float64 `json:"measureValue"`
	Timestamp    int     `json:"timestamp"`
}
