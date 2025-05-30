package models

import "time"

type OrderItem struct {
	ID        uint64    `gorm:"primary_key,auto_increment" json:"id"`
	OrderID   uint64    `gorm:"not_null" json:"orderId"`
	ShopID    uint64    `gorm:"not_null" json:"shopId"`
	ProductID uint64    `gorm:"not_null" json:"ProductId"`
	Quantity  uint8     `gorm:"not_null" json:"quantity"`
	Price     float32   `gorm:"not_null" json:"price"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
