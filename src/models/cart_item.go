package models

import "time"

type CartItem struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CartID    uint64    `gorm:"not_null" json:"cartId"`
	ProductID uint64    `gorm:"not_null" json:"productId"`
	Quantity  uint8     `gorm:"not_null" json:"quantity"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
