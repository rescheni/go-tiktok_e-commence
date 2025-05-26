package mysql

import (
	"e-commence/app/order/biz/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"k8s.io/klog/v2"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
	DB, err = gorm.Open(mysql.Open(os.Getenv("GOMALL_MYSQL_DSN")),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	err := DB.AutoMigrate(
		&model.Order{},
		&model.OrderItem{},
	)
	if err != nil {
		klog.Error(err)
	}
}
