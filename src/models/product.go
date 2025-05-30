package models

import "time"

type Product struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	ShopID      uint64    `gorm:"not_null" json:"shopId"`
	Name        string    `gorm:"not_null" json:"name"`
	Description string    `gorm:"not_null" json:"description"`
	Price       float32   `gorm:"not_null" json:"price"`
	Stock       uint16    `gorm:"not_null" json:"stock"`
	Category    string    `gorm:"not_null" json:"category"`
	SubCategory string    `gorm:"not_null" json:"subCategory"`
	Popularity  uint8     `gorm:"not_null" json:"popularity"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
