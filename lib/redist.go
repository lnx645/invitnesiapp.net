package lib

import (
	"context"
	"fmt"
	"invitnesia/api/config"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedisClient(cfg *config.Redis) {
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
	if err != nil {
		panic("Redis Server Tidak Aktif")
	}
	fmt.Println("✅ Berhasil terhubung ke Memurai!")
}
