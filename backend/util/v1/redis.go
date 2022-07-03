package v1

import (
	"fmt"

	"github.com/go-redis/redis"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "kump-redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
