package dal

import (
	"e-commence/app/order/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
