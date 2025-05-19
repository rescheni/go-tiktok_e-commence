package redis

import (
	"context"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func Init() {
	// RedisClient = redis.NewClient(&redis.Options{
	// 	Addr:     conf.GetConf().Redis.Address,
	// 	Username: conf.GetConf().Redis.Username,
	// 	Password: conf.GetConf().Redis.Password,
	// 	DB:       conf.GetConf().Redis.DB,
	// })
	dbid, err := strconv.Atoi(os.Getenv("GOMALL_REDIS_DB"))
	if err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("GOMALL_REDIS_ADDRESS"),
		Username: os.Getenv("GOMALL_REDIS_USERNAME"),
		Password: os.Getenv("GOMALL_REDIS_PASSWORD"),
		DB:       dbid,
	})

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
