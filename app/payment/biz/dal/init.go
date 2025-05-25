package dal

import "e-commence/app/payment/biz/dal/mysql"

func Init() {
	// redis.Init()
	mysql.Init()
}
