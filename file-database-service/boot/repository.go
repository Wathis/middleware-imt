package boot

import (
	"file-database-service/application"
	"file-database-service/infrastructure/repository"
)

func LoadRepositories() {
	// Initialisation des connecteurs vers les interfaces (DAO)
	application.MeasureRepository = repository.NewMeasureRepository()
}
