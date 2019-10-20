package application

import (
	"runtime"
)

func LaunchRedisServer() {
	if runtime.GOOS == "windows" {
		executeWindows()
	} else {
		execute()
	}
}

func execute() {
	// TODO : Lancer le serveur redis
}

func executeWindows() {
	// TODO : Lancer le serveur redis
}
