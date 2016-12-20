package redis

import (
	"time"
)

func Get(key string) (string, error) {
	val, err := client.Get(key).Result()

	return val, err
}

func Set(key string, value string, expiration time.Duration) error {
	return client.Set(key, value, expiration).Err()
}

func Delete(key string) (int64, error) {
	val, err := client.Del(key).Result()

	return val, err
}

func Increment(key string) (int64, error) {
	return client.Incr(key).Result()
}

func Expire(key string, expiration time.Duration) (bool, error) {
	return client.Expire(key, expiration).Result()
}
