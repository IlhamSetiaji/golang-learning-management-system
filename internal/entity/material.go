package entity

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	MasterClassID       uint                 `json:"master_class_id"`
	MasterClass         MasterClass          `gorm:"foreignKey:MasterClassID;references:ID"`
	Title               string               `json:"title"`
	Description         string               `json:"description" gorm:"type:text"`
	MaterialAttachments []MaterialAttachment `gorm:"foreignKey:MaterialID;references:ID"`
}
