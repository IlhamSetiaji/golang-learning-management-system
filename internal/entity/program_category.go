package entity

import "gorm.io/gorm"

type ProgramCategory struct {
	gorm.Model
	Name     string    `json:"name" gorm:"unique"`
	Programs []Program `gorm:"foreignKey:ProgramCategoryID;references:ID"`
}
