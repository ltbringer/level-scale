package models

import "time"

type User struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	IsSeller     bool      `gorm:"default:false" json:"isSeller"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
