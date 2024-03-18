package entity

import "gorm.io/gorm"

type Timeline struct {
	gorm.Model
	MasterClassID   uint             `json:"master_class_id"`
	MasterClass     MasterClass      `gorm:"foreignKey:MasterClassID;references:ID"`
	Title           string           `json:"title"`
	Target          string           `json:"target"`
	TimelineDetails []TimelineDetail `gorm:"foreignKey:TimelineID;references:ID"`
}
