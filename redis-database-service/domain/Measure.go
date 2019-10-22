package entities

// Mesure : Structure stockant les informations envoyées en json pour décrire une mesure
type Measure struct {
	SensorID     int     `json:"sensorId"`
	AirportID    string  `json:"airportId"`
	MeasureType  string  `json:"measureType"`
	MeasureValue float64 `json:"measureValue"`
	Timestamp    int     `json:"timestamp"`
}
