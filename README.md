# middleware-imt

## sensors-service : 

### config file : 

```json
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

### PUBLISH sensor/mesure :

```json
{
  "sensorId" : int,
  "airportId" : string (IATA), 
  "mesureType" : string,
  "mesureValue" : float,
  "timestamp": unix timestamp (seconds)
}
```

## redis-database-service : 

### SUBSCRIBER sensor/mesure

## file-database-service : 

### SUBSCRIBER sensor/mesure





