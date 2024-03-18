package entity

import "gorm.io/gorm"

type MenteeTask struct {
	gorm.Model
	UserID                uint                   `json:"user_id"`
	User                  User                   `gorm:"foreignKey:UserID;references:ID"`
	TaskID                uint                   `json:"task_id"`
	Task                  Task                   `gorm:"foreignKey:TaskID;references:ID"`
	Description           string                 `json:"description" gorm:"type:text"`
	Score                 float64                `json:"score"`
	Comment               string                 `json:"comment" gorm:"type:text"`
	Status                string                 `json:"status"`
	MenteeTaskAttachments []MenteeTaskAttachment `gorm:"foreignKey:MenteeTaskID;references:ID"`
}
