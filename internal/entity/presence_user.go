package entity

import "time"

type PresenceUser struct {
	PresenceID uint      `gorm:"primaryKey" json:"presence_id"`
	UserID     uint      `gorm:"primaryKey" json:"user_id"`
	Summary    string    `json:"summary" gorm:"type:text"`
	Duration   int       `json:"duration"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
