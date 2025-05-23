package dal

import (
	"e-commence/app/cart/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
