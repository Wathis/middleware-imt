package domain

// Measure : Structure stockant les informations envoyées en json pour décrire une mesure
type Measure struct {
	SensorID     int64   `json:"sensorId"`
	AirportID    string  `json:"airportId"`
	MeasureType  string  `json:"measureType"`
	MeasureValue float64 `json:"measureValue"`
	Timestamp    int64   `json:"timestamp"`
}
