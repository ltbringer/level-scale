package models

import "time"

type Order struct {
	ID                             uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID                         uint64    `gorm:"not_null" json:"userId"`
	CartID                         uint64    `gorm:"not_null" json:"cartId"`
	ShippingAddress                string    `gorm:"not_null" json:"shippingAddress"`
	BillingAddress                 string    `json:"billingAddress"`
	IsShippingAddressSameAsBilling string    `gorm:"default:false" json:"isShippingAddressSameAsBilling"`
	CardNumber                     string    `gorm:"not_null" json:"-"`
	Status                         string    `gorm:"not_null" json:"status"`
	CreatedAt                      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
