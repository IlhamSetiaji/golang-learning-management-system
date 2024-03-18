package entity

import (
	"time"

	"gorm.io/gorm"
)

type Presence struct {
	gorm.Model
	UserID        uint        `json:"user_id"`
	User          User        `gorm:"foreignKey:UserID;references:ID"`
	MasterClassID uint        `json:"master_class_id"`
	MasterClass   MasterClass `gorm:"foreignKey:MasterClassID;references:ID"`
	Name          string      `json:"name"`
	Summary       string      `json:"summary" gorm:"type:text"`
	Duration      int         `json:"duration"`
	StartTime     time.Time   `json:"start_time"`
	EndTime       time.Time   `json:"end_time"`
	Users         []User      `gorm:"many2many:presence_users;"`
}
