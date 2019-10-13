package main

import (
	"api-service/boot"
	"log"
)

// @title Swagger API documentation for api-service
// @version 1.0
// @license.name GNU GPLv3
// @license.url https://www.gnu.org/licenses/gpl-3.0.fr.html
func main() {
	log.Print("starting api-service")
	log.Print("Loading env")
	boot.LoadEnv()
	log.Print("Loading repositories")
	boot.LoadRepositories()
	log.Print("Loading http router")
	boot.LoadRouter()
}
