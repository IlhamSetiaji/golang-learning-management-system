package entity

import "gorm.io/gorm"

type Program struct {
	gorm.Model
	ParentID          uint            `json:"parent_id" gorm:"default:null"` // 0 for no parent
	Parent            *Program        `gorm:"foreignKey:ParentID;references:ID"`
	ProgramCategoryID uint            `json:"program_category_id"`
	ProgramCategory   ProgramCategory `gorm:"foreignKey:ProgramCategoryID;references:ID"`
	Name              string          `json:"name"`
	Description       string          `gorm:"type:text" json:"description"`
	EnrollCode        string          `json:"enroll_code" gorm:"unique"`
	Price             float64         `json:"price"` // Exported field
	Duration          int             `json:"duration"`
	Level             string          `json:"level"`
	IsActive          bool            `json:"is_active" gorm:"default:true"`
	Batches           []Batch         `gorm:"foreignKey:ProgramID;references:ID"`
}
