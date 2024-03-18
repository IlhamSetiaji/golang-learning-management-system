package entity

import "time"

type PresenceUser struct {
	PresenceID uint      `gorm:"primaryKey" json:"presence_id"`
	UserID     uint      `gorm:"primaryKey" json:"user_id"`
	Summary    string    `json:"summary" gorm:"type:text"`
	Duration   int       `json:"duration"`
	StartTime  time.Time `json:"start_time" gorm:"type:datetime;"` // Added closing double quote
	EndTime    time.Time `json:"end_time" gorm:"type:datetime;"`   // Added closing double quote
	CreatedAt  time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
