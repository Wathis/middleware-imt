package repository

import (
	"encoding/csv"
	"file-database-service/domain"
	"file-database-service/internal/exportpath"
	"log"
	"os"
	"strconv"
	"time"
)

type MeasureRepository struct {
}

func NewMeasureRepository() *MeasureRepository {
	return &MeasureRepository{}
}

// SaveMeasure save all the data in the database
func (m *MeasureRepository) SaveMeasure(measure domain.Measure) error {
	err := saveMeasureToFile(measure)
	return err
}

func saveMeasureToFile(measure domain.Measure) error {

	var measureDatetime string = time.Unix(measure.Timestamp, 0).Format("2006-01-02 15:04:05")
	var filename string = string(measure.AirportID) + "-" + measureDatetime[:10] + "-" + string(measure.MeasureType) + ".csv"

	_ = os.Mkdir(exportpath.FileExportPath, 0755) // cree le dossier si il n'existe pas déjà
	f, err := os.OpenFile(exportpath.FileExportPath+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	// Write line in file
	writer.Write([]string{
		strconv.FormatInt(measure.SensorID, 10),
		measure.AirportID,
		measure.MeasureType,
		strconv.FormatFloat(measure.MeasureValue, 'f', 6, 32),
		measureDatetime[:19],
	})

	log.Print("Mesure sauvegardee dans le fichier: ", exportpath.FileExportPath+"/"+filename)
	return nil
}
