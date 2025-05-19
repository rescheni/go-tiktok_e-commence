package dal

import (
	"e-commence/app/user/biz/dal/mysql"
	"e-commence/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
