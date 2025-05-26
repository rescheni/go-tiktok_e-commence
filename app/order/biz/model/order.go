package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	Ciry          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId       uint64      `gorm:"type:int(11);"`
	UserCurrency string      `gorm:"type:varchar(10)"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItem    []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

// TODO: 单元测试
func (o Order) TableName() string {
	return "order"
}
func ListOrder(ctx context.Context, DB *gorm.DB, user_id uint64) ([]*Order, error) {

	var orders []*Order
	err := DB.WithContext(ctx).Where("user_id = ?", user_id).Preload("OrderItem").Find(&orders).Error

	if err != nil {
		return nil, err
	}
	return orders, nil
}
