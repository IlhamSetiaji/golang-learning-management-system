package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	MasterClassID   uint             `json:"master_class_id"`
	MasterClass     MasterClass      `gorm:"foreignKey:MasterClassID;references:ID"`
	Title           string           `json:"title"`
	Description     string           `json:"description" gorm:"type:text"`
	StartTime       time.Time        `json:"start_time" gorm:"type:datetime;`
	EndTime         time.Time        `json:"end_time" gorm:"type:datetime;`
	IsDone          bool             `json:"is_done" gorm:"default:false"`
	MenteeTasks     []MenteeTask     `gorm:"foreignKey:TaskID;references:ID"`
	TaskAttachments []TaskAttachment `gorm:"foreignKey:TaskID;references:ID"`
}
