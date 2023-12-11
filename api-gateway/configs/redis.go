package configs

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	rHost := os.Getenv("REDIS_HOST")
	rPort := os.Getenv("REDIS_PORT")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rHost, rPort),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	return client
}
