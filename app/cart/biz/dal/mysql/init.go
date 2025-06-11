package mysql

import (
	"e-commence/app/cart/biz/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
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

	DB.AutoMigrate(
		&model.Cart{},
	)
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

}
