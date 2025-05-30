package models

import "time"

type Shop struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	SellerID    uint64    `gorm:"not_null" json:"sellerId"`
	Name        string    `gorm:"not_null" json:"name"`
	Description string    `gorm:"not_null" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
