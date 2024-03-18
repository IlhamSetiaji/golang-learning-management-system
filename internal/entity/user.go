package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uint          `json:"id" gorm:"primary_key"` // gorm.Model is a struct that contains ID, CreatedAt, UpdatedAt, DeletedAt
	Name            string        `json:"name"`
	Email           string        `json:"email" gorm:"unique"`
	Username        string        `json:"username" gorm:"unique"`
	Password        string        `json:"password"`
	EmailVerifiedAt time.Time     `json:"email_verified_at" gorm:"default:null"`
	IsActive        bool          `json:"is_active" gorm:"default:true"`
	Roles           []Role        `gorm:"many2many:user_roles;"`
	MasterClasses   []MasterClass `gorm:"many2many:master_class_users;"`
	Presences       []Presence    `gorm:"foreignKey:UserID;references:ID"`
	PresenceUsers   []Presence    `gorm:"many2many:presence_users;"`
	MenteeTasks     []MenteeTask  `gorm:"foreignKey:UserID;references:ID"`
}
