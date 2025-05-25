package dal

import (
	"e-commence/app/checkout/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
