package models

import "time"

type Delivery struct {
	ID         uint64    `gorm:"primaryKey;auto_increment" json:"id"`
	OrderID    uint64    `gorm:"not null" json:"orderId"`
	ExpectedAt time.Time `gorm:"not null" json:"scheduledFor"`
	Status     string    `gorm:"default:'scheduled'" json:"status"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
