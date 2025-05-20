package dal

import (
	"e-commence/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
