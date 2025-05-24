package dal

import (
	"e-commence/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	// mysql.Init()
}
