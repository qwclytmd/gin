package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func ConnectionRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     v.GetString("redis.addr"),
		Password: v.GetString("redis.password"),
		DB:       v.GetInt("redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		//log.Fatalln("Redis connection error: ", err)
	}

	log.Println("Redis connection success")

	return client
}
