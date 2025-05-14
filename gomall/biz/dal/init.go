package dal

import (
	"e-commence/gomall/biz/dal/mysql"
	"e-commence/gomall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
