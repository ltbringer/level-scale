package models

import "time"

type Invoice struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	OrderID   uint64    `gorm:"not_null" json:"orderId"`
	Amount    float32   `gorm:"not_null" json:"amount"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
