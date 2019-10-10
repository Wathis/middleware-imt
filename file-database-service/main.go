package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type sensorData struct {
	SensorID    int64
	AirportID   string
	MesureType  string
	MesureValue float64
	Timestamp   int64
}

const testData = `{"SensorID": 43, "AirportID": "NTE", "mesureType": "temp", "mesureValue": 17.33, "timestamp": 1570731034}`

func main() {

	// Each time we receive data from broker
	data := sensorData{}
	json.Unmarshal([]byte(testData), &data)
	fmt.Println(data)
	addDataToCsv(data)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func addDataToCsv(row sensorData) {
	rowDatetime := getTime(row.Timestamp)

	var filename string = string(row.AirportID) + "-" + rowDatetime[:10] + "-" + string(row.MesureType) + ".csv"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(f)
	defer writer.Flush()
	writer.Write([]string{
		strconv.FormatInt(row.SensorID, 10),
		row.AirportID,
		row.MesureType,
		strconv.FormatFloat(row.MesureValue, 'f', 6, 32),
		rowDatetime,
	})
	checkError("Cannot write to file", err)
}

func getTime(input int64) string {
	var tm time.Time = time.Unix(input, 0)
	var res string = tm.Format("2006-01-02 15:04:05")
	return res
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
