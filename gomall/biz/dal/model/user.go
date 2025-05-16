package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(128);unique"`
	Email    string `gorm:"type:varchar(128);unique_index"`
	Password string `gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "user"
}
