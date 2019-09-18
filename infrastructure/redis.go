package infrastructure

import (
	"fmt"

	"github.com/go-redis/redis"
)

// RedisDB to connect to Redis
func RedisDB() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       9,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("Error connecting to Redis: %s", err)
	}

	return client
}
