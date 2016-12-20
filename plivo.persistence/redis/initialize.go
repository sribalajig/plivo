package redis

import (
	"gopkg.in/redis.v3"
)

var client *redis.Client

func Initialize() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
