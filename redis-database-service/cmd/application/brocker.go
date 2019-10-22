package application

import (
	"encoding/json"
	"fmt"
	"log"
	"redis-database-service/cmd/boot"
	"redis-database-service/cmd/dao"
	"redis-database-service/cmd/redis"
	entities "redis-database-service/internal/entities"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	MqttClient *mqtt.Client
)
