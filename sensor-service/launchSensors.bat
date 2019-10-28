REM CDG
START sensors/sensorService.exe sensors/configSensorTempCDG.json
START sensors/sensorService.exe sensors/configSensorWindCDG.json
START sensors/sensorService.exe sensors/configSensorPressureCDG.json

timeout /t 2 /nobreak > nul

REM LYS
START sensors/sensorService.exe sensors/configSensorTempLYS.json
START sensors/sensorService.exe sensors/configSensorWindLYS.json
START sensors/sensorService.exe sensors/configSensorPressureLYS.json

timeout /t 2 /nobreak > nul

REM NTE
START sensors/sensorService.exe sensors/configSensorTempNTE.json
START sensors/sensorService.exe sensors/configSensorWindNTE.json
START sensors/sensorService.exe sensors/configSensorPressureNTE.json