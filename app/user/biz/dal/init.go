package dal

import (
	"e-commence/app/user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
