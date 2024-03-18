package entity

import (
	"time"

	"gorm.io/gorm"
)

type Batch struct {
	gorm.Model
	ProgramID             uint          `json:"program_id"`
	Program               Program       `gorm:"foreignKey:ProgramID;references:ID"`
	Name                  string        `json:"name"`
	RegistrationStartDate time.Time     `json:"registration_start_date"`
	RegistrationEndDate   time.Time     `json:"registration_end_date"`
	StartDate             time.Time     `json:"start_date"`
	EndDate               time.Time     `json:"end_date"`
	MasterClasses         []MasterClass `gorm:"foreignKey:BatchID;references:ID"`
}
