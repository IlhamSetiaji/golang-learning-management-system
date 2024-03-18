package entity

import "gorm.io/gorm"

type TaskAttachment struct {
	gorm.Model
	TaskID uint   `json:"task_id"`
	Task   Task   `gorm:"foreignKey:TaskID;references:ID"`
	Name   string `json:"name"`
	Path   string `json:"path" gorm:"type:text"`
	Type   string `json:"type"`
}
