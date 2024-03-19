package entity

import "gorm.io/gorm"

type Program struct {
	gorm.Model
	ProgramCategoryID uint            `json:"program_category_id"`
	ProgramCategory   ProgramCategory `gorm:"foreignKey:ProgramCategoryID;references:ID"`
	Name              string          `json:"name"`
	Description       string          `gorm:"type:text" json:"description"`
	Image             string          `json:"image" gorm:"type:text;default:null"`
	IsActive          bool            `json:"is_active" gorm:"default:true"`
	Batches           []Batch         `gorm:"foreignKey:ProgramID;references:ID"`
}
