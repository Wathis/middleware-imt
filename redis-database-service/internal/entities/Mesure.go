package entities

// Mesure : Structure stockant les informations envoyées en json pour décrire une mesure
type Mesure struct {
	SensorID    int     `json:"sensorId"`
	AirportID   string  `json:"airportId"`
	MesureType  string  `json:"measureType"`
	MesureValue float64 `json:"measureValue"`
	Timestamp   int     `json:"timestamp"`
}
