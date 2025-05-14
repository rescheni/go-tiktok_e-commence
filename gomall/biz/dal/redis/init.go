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
	//RedisClient = redis.NewClient(&redis.Options{
	//	Addr:     conf.GetConf().Redis.Address,
	//	Username: conf.GetConf().Redis.Username,
	//	Password: conf.GetConf().Redis.Password,
	//	DB:       conf.GetConf().Redis.DB,
	//})
	db, err := strconv.Atoi(os.Getenv("GOMALL_REDIS_DB"))
	if err != nil {
		return
	}

	RedisClient = redis.NewClient(&redis.Options{
		//Addr:     conf.GetConf().Redis.Address,
		Addr: os.Getenv("GOMALL_REDIS_URL" + ":" + os.Getenv("GOMALL_REDIS_PORT")),
		//Username: conf.GetConf().Redis.Username,
		Username: "",
		//Password: conf.GetConf().Redis.Password,
		Password: "",
		//DB:       conf.GetConf().Redis.DB,
		DB: db,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
