package entity

import "gorm.io/gorm"

type Objective struct {
	gorm.Model
	MasterClassID uint        `json:"master_class_id"`
	MasterClass   MasterClass `gorm:"foreignKey:MasterClassID;references:ID"`
	Title         string      `json:"title"`
	Description   string      `json:"description" gorm:"type:text"`
}
