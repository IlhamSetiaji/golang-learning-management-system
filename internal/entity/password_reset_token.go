package entity

import "time"

// PasswordResetToken is a struct that represent the password_reset_tokens table in the database
type PasswordResetToken struct {
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
