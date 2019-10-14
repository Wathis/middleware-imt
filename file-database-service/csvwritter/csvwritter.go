package csvwritter

import (
	"encoding/csv"
	"file-database-service/sensor"
	"log"
	"os"
	"strconv"
	"time"
)

func AddDataToCsv(row sensor.SensorData, filePath string) {
	rowDatetime := getTime(row.Timestamp)

	var filename string = string(row.AirportID) + "-" + rowDatetime[:10] + "-" + string(row.MesureType) + ".csv"

	f, err := os.OpenFile(filePath+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	// Write line here
	writer.Write([]string{
		strconv.FormatInt(row.SensorID, 10),
		row.AirportID,
		row.MesureType,
		strconv.FormatFloat(row.MesureValue, 'f', 6, 32),
		rowDatetime,
	})

	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
}

func getTime(input int64) string {
	var tm time.Time = time.Unix(input, 0)
	var res string = tm.Format("2006-01-02 15:04:05")
	return res
}
