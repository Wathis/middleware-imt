# middleware-imt

## sensors-service

### config file 

```
{
    "intervalSendDataSensor" : int64,
    "brokerUrl" : string,
    "brokerPort" : int64,
    "sensorId" : int64,
    "sensorType" : string,
    "qos" : byte,
    "min" : float64,
    "max" : float64,
    "variation" : float64
}
```

### PUBLISH on topic sensor/measure

```
{
  "sensorId" : int64,
  "airportId" : string (IATA), 
  "measureType" : string,
  "measureValue" : float64,
  "timestamp": int64
}
```

## redis-database-service 

### SUBSCRIBE on topic sensor/measure
```
{
  "sensorId" : int,
  "airportId" : string (IATA), 
  "measureType" : string,
  "measureValue" : float,
  "timestamp": unix timestamp (seconds)
}
```

### data storage 

[Redis scheme](https://github.com/Wathis/middleware-imt/blob/develop/docs/redis_scheme.png)

#### Test data

```
MULTI
SET sensor:1:measure:1570966444 '{"sensorId":1,"airportId":"CDG","measureType":"temperature","measureValue":10.7,"timestamp":1570966444}'
ZADD measure_value:temperature 10.7 sensor:1:measure:1570966444
ZADD measure_timestamp:temperature 1570966444 sensor:1:measure:1570966444
SET sensor:1:measure:1570966450 '{"sensorId":1,"airportId":"CDG","measureType":"temperature","measureValue":10.4,"timestamp":1570966450}'
ZADD measure_value:temperature 10.4 sensor:1:measure:1570966450
ZADD measure_timestamp:temperature 1570966450 sensor:1:measure:1570966450
EXEC
```


## file-database-service 

### SUBSCRIBE on topic sensor/measure
```
{
  "sensorId" : int,
  "airportId" : string (IATA), 
  "measureType" : string,
  "measureValue" : float,
  "timestamp": unix timestamp (seconds)
}
```


## api-service 

Swagger api with : https://editor.swagger.io

Swagger xml doc : [swagger yml doc](https://github.com/Wathis/middleware-imt/blob/develop/api-service/docs/swagger.yaml)

Just copy and paste xml doc into swagger editor



