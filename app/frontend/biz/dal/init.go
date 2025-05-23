package dal

import (
	"gomall/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
