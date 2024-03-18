package entity

import "gorm.io/gorm"

type MenteeTaskAttachment struct {
	gorm.Model
	MenteeTaskID uint       `json:"mentee_task_id"`
	MenteeTask   MenteeTask `gorm:"foreignKey:MenteeTaskID;references:ID"`
	Name         string     `json:"name"`
	Path         string     `json:"path" gorm:"type:text"`
	Type         string     `json:"type"`
}
