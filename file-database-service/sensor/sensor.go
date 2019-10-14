package sensor

type SensorData struct {
	SensorID    int64
	AirportID   string
	MesureType  string
	MesureValue float64
	Timestamp   int64
}
