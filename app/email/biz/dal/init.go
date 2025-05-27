package dal

import (
	"e-commence/app/email/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
