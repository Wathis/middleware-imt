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

### PUBLISH on topic sensor/measure

```
{
  "sensorId" : int,
  "airportId" : string (IATA), 
  "measureType" : string,
  "measureValue" : float,
  "timestamp": unix timestamp (seconds)
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

```sensor:<sensorId>:measure:<timestamp> <json>```


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


### GET /sensors
```
[
  {
    "sensorId" : int,
    "airportId" : string (IATA), 
  },
  ...
]
```

### GET /sensors/{sensorId}
```
{
  "sensorId" : int,
  "airportId" : string (IATA), 
  "measureType" : string
}
```

### GET /sensors/{sensorId}/measures
```
[
  {
    "sensorId" : int,
    "airportId" : string (IATA), 
    "measureType" : string,
    "measureValue" : float,
    "timestamp": unix timestamp (seconds) 
  },
  ...
]
```

### GET /mesures
```
[
  {
    "sensorId" : int,
    "airportId" : string (IATA), 
    "measureType" : string,
    "measureValue" : float,
    "timestamp": unix timestamp (seconds)
  },
  ...
]
```




