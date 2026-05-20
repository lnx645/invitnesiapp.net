package lib

import (
	"context"
	"fmt"
	"invitnesia/api/app/config"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// InitRedisClient initializes the Redis client with the provided configuration.
func InitRedisClient(cfg *config.Redis) error {
	dat, err := strconv.Atoi(cfg.DB)
	if err != nil {
		panic(err)
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       dat,
	})

	context := context.Background()
	_, err = RedisClient.Ping(context).Result()
	return err
}
