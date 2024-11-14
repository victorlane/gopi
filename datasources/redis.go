package datasources

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func InitRedis(password string, host string, port string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	return client
}
