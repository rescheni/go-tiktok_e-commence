package dal

import (
	"e-commence/app/frontend/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
