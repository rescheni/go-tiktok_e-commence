package dal

import (
	"e-commence/biz/dal/mysql"
	"e-commence/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
