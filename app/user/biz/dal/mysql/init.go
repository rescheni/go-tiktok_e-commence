package mysql

import (
	"e-commence/app/user/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	if DB == nil {
		panic("mysql db is nil")
	}

	DB.AutoMigrate(
		&model.User{},
	)

}
