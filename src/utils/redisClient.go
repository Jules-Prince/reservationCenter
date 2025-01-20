package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Default Redis address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})

	// Test the connection
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully!")
}
