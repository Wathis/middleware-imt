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

}

func executeWindows() {

}
