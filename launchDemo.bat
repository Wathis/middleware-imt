START file-database-service/main.exe
START api-service/main.exe
START redis-database-service/main.exe

START sensor-service/sensors/sensorService.exe   sensor-service/sensors/configSensorPressureCDG.json
START sensor-service/sensors/sensorService.exe   sensor-service/sensors/configSensorTempCDG.json
START sensor-service/sensors/sensorService.exe   sensor-service/sensors/configSensorWindCDG.json

START sensor-service/sensors/sensorService.exe   sensor-service/sensors/configSensorPressureNTE.json
START sensor-service/sensors/sensorService.exe   sensor-service/sensors/configSensorTempNTE.json
START sensor-service/sensors/sensorService.exe   sensor-service/sensors/configSensorWindNTE.json

START sensor-service/sensors/sensorService.exe   sensor-service/sensors/configSensorPressureLYS.json
START sensor-service/sensors/sensorService.exe   sensor-service/sensors/configSensorTempLYS.json
START sensor-service/sensors/sensorService.exe   sensor-service/sensors/configSensorWindLYS.json