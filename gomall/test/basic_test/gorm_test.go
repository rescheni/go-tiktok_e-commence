package basic

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type test_table struct {
	//Id          int64     `gorm:"primary_key"`
	Create_time time.Time `gorm:"type:datetime;not null"`
	Key         string
	Value       string
}

func gorm_init() *gorm.DB {

	err := godotenv.Load("../../.env")
	if err != nil {
	}

	DB, err := gorm.Open(mysql.Open(os.Getenv("GOMALL_MYSQL_DSN")),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	return DB
}

func TestInsertGorm(t *testing.T) {

	db := gorm_init()
	err := db.AutoMigrate(&test_table{})
	if err != nil {
		return
	}
	log.Println("db.AutoMigrate OK")
	db.Create(&test_table{
		Create_time: time.Now(),
		Key:         "key1",
		Value:       "value1",
	})
	db.Create([]test_table{
		{
			Create_time: time.Now(),
			Key:         "key2",
			Value:       "value2",
		},
		{
			Create_time: time.Now(),
			Key:         "key3",
			Value:       "value3",
		},
	})
}

func TestCRUDGorm(t *testing.T) {
	db := gorm_init()
	cnt := int64(0)
	// 增
	db.Create(&test_table{
		Create_time: time.Now(),
		Key:         "key0",
		Value:       "value0",
	})
	db.Where("key = ?", "key0").Count(&cnt)
	log.Println(cnt)
	// 改
	db.Model(&test_table{}).Where("`key` = ?", "key0").Update("value", "value_update")
	// 查找
	var get_db test_table
	db.Model(&test_table{}).Where("`key` = ?", "key0").First(&get_db)
	log.Println(get_db.Key, get_db.Value, get_db.Create_time)
	// 删
	db.Where("`key` = ?", "key0").Delete(&test_table{})

	db.Where("`key` = ?", "key0").Count(&cnt)
	log.Println(cnt)
}

func TestGetGorm(b *testing.T) {
	db := gorm_init()
	tables := &[]test_table{}
	result := db.Find(&tables)
	if result.Error != nil {
		b.Fatal(result.Error)
	}
	for _, table := range *tables {
		fmt.Println(table.Create_time, table.Key, table.Value)
	}
}
