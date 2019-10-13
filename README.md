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

Swagger api with : https://editor.swagger.io

Swagger xml doc : [swagger yml doc](https://github.com/Wathis/middleware-imt/blob/develop/api-service/docs/swagger.yaml)

Just copy and paste xml doc into swagger editor



