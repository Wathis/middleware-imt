package dao

import (
	redis "github.com/go-redis/redis/v7"
)

func InitDatabase(host string) {
	Connexion = connecting(host)
}

// Connecting : Retourne une structure contenant les infos de connection à la base de données
func connecting(host string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: host,
	})
}
