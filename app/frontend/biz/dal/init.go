package dal

import (
	"e-commence/app/user/biz/dal/redis"
	"gomall/biz/dal/mysql"
)

func Init() {
	redis.Init()
	mysql.Init()
}
