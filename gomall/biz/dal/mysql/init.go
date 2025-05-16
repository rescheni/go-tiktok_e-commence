package mysql

import (
	"e-commence/gomall/biz/dal/model"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := os.Getenv("GOMALL_MYSQL_DSN")
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", DB.Debug().Exec("SHOW TABLES"))
}
