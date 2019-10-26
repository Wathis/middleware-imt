module sensorPressure

require (
	common v0.0.0
	config v0.0.0
	mqtt v0.0.0
)

replace common => ../common

replace config => ../config

replace mqtt => ../mqtt

go 1.13
