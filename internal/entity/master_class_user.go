package entity

import "time"

type MasterClassUser struct {
	MasterClassID uint      `gorm:"primaryKey" json:"master_class_id"`
	UserID        uint      `gorm:"primaryKey" json:"user_id"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
