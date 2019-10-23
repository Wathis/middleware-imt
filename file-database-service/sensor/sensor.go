package sensor

type SensorData struct {
	SensorID     int64
	AirportID    string
	MeasureType  string
	MeasureValue float64
	Timestamp    int64
}
