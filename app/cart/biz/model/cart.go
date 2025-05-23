package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint64 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint64 `gorm:"type:int(11);not null;"`
	Qty       uint32 `gorm:"type:int(11);not null;"`
}

func (c Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, cart Cart) (err error) {

	var row Cart

	// 查询当前商品是否已经存在于购物车 ？ 如果已经存在
	err = db.WithContext(ctx).Model(&Cart{}).
		Where("user_id = ? AND product_id = ?", cart.UserId, cart.ProductId).
		First(&row).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: cart.UserId, ProductId: cart.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", cart.Qty)).
			Error
	}

	return db.WithContext(ctx).Create(&cart).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, user_id uint64) (err error) {

	if user_id == 0 {
		return errors.New("userid is empty")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", user_id).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, user_id uint64) (carts []*Cart, err error) {

	err = db.WithContext(ctx).
		Model(&Cart{}).
		Where("user_id = ?", user_id).
		Find(&carts).
		Error
	return
}
