package entity

import (
	"time"
)

type UserRole struct {
	UserID    uint      `gorm:"primaryKey" json:"user_id"`
	RoleID    uint      `gorm:"primaryKey" json:"role_id"`
	CreatedAt time.Time `type:datetime;gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `type:datetime;gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
