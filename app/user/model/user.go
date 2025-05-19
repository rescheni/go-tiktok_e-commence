package model

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"unique;not null;type:varchar(128);comment:用户名"`
	PasswordHashed string `json:"password" gorm:"not null;type:varchar(255);comment:密码"`
	Email          string `json:"email" gorm:"unique; not null;type:varchar(255);comment:邮箱"`
}

func (User) TableName() string {
	return "user"
}
func Create(db *gorm.DB, user *User) error {

	if err := db.Create(user).Error; err != nil {
		klog.Errorf("Create user failed, err: %v", err)
	}
	return nil
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		klog.Errorf("Get user by username failed, err: %v", err)
		return nil, err
	}
	return &user, nil
}
func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		klog.Errorf("Get user by email failed, err: %v", err)
		return nil, err
	}
	return &user, nil
}
