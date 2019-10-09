# middleware-imt

## sensors-service

### config file 

```
[
  {
    "brokerUrl" : string,
    "borkerPort" : int,
    "sensorId" : int,
    "sensorType" : string,
    "qos" : int
  },
  ...
]
```

### PUBLISH on topic sensor/mesure

```
{
  "sensorId" : int,
  "airportId" : string (IATA), 
  "mesureType" : string,
  "mesureValue" : float,
  "timestamp": unix timestamp (seconds)
}
```

## redis-database-service 

### SUBSCRIBE on topic sensor/mesure
```
{
  "sensorId" : int,
  "airportId" : string (IATA), 
  "mesureType" : string,
  "mesureValue" : float,
  "timestamp": unix timestamp (seconds)
}
```

### data storage 

```sensor:<sensorId>:mesure:<timestamp> <json>```


## file-database-service 

### SUBSCRIBE on topic sensor/mesure
```json
{
  "sensorId" : int,
  "airportId" : string (IATA), 
  "mesureType" : string,
  "mesureValue" : float,
  "timestamp": unix timestamp (seconds)
}
```





