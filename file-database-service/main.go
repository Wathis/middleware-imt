package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type sensorData struct {
	SensorID    int64
	AirportID   string
	MesureType  string
	MesureValue float64
	Timestamp   float64
}

const testData = `{"SensorID": 0, "AirportID": "NTE", "mesureType": "temp", "mesureValue": 17.33, "timestamp": 1570652403}`

func timestampToDate(val interface{}) string {
	return "2019-11-10:14:45:10"
}

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
	var filename string = string(row.AirportID) + "-" + timestampToDate(row.Timestamp)[:10] + "-" + string(row.MesureType) + ".csv"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(f)
	defer writer.Flush()
	writer.Write([]string{strconv.FormatInt(row.SensorID, 32), row.AirportID, row.MesureType, floatToString(row.MesureValue), floatToString(row.Timestamp)})
	checkError("Cannot write to file", err)
}

func floatToString(val float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(val, 'f', 6, 32)
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
