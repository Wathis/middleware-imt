package dao

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	Connexion *redis.Pool
)

func InitDatabase(host string) {
	Connexion = connecting(host)
}

// Connecting : Retourne une structure contenant les infos de connection à la base de données
func connecting(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host)
		},
	}
}
