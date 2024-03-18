package entity

import "gorm.io/gorm"

type MaterialAttachment struct {
	gorm.Model
	MaterialID uint     `json:"material_id"`
	Material   Material `gorm:"foreignKey:MaterialID;references:ID"`
	Name       string   `json:"name"`
	Path       string   `json:"path" gorm:"type:text"`
	Type       string   `json:"type"`
}
