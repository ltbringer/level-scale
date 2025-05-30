package models

import "time"

type Return struct {
	ID           uint      `gorm:"primaryKey;auto_increment" json:"id"`
	OrderItemID  uint      `gorm:"not null" json:"orderItemId"`
	ReturnReason string    `gorm:"not null" json:"reason"`
	Status       string    `gorm:"default:'requested'" json:"status"`
	RejectReason string    `gorm:"not null" json:"rejectReason"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
