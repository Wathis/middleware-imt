# middleware-imt

## sensors-service : 

### config file : 

``
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
``

### PUBLISH sensor/mesure :

`
{
  "sensorId" : int,
  "airportId" : string (IATA), 
  "mesureType" : string,
  "mesureValue" : float,
  "timestamp": unix timestamp (seconds)
}
`

## redis-database-service : 

### SUBSCRIBER sensor/mesure

## file-database-service : 

### SUBSCRIBER sensor/mesure





