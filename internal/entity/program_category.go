package entity

import "gorm.io/gorm"

type ProgramCategory struct {
	gorm.Model
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description" gorm:"type:text"`
	Image       string    `json:"image" gorm:"type:text;default:null"`
	Programs    []Program `gorm:"foreignKey:ProgramCategoryID;references:ID"`
}
