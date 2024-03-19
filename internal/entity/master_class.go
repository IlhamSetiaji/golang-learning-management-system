package entity

import "gorm.io/gorm"

type MasterClass struct {
	gorm.Model
	ParentID    uint          `json:"parent_id" gorm:"default:null"` // 0 for no parent
	Parent      *MasterClass  `gorm:"foreignKey:ParentID;references:ID"`
	Children    []MasterClass `gorm:"foreignKey:ParentID;references:ID"`
	BatchID     uint          `json:"batch_id"`
	Batch       Batch         `gorm:"foreignKey:BatchID;references:ID"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Description string        `json:"description"`
	Image       string        `json:"image" gorm:"type:text;default:null"`
	EnrollCode  string        `json:"enroll_code" gorm:"unique;default:null"`
	Price       float64       `json:"price"` // Exported field
	Duration    int           `json:"duration"`
	Level       string        `json:"level"`
	IsActive    bool          `json:"is_active" gorm:"default:true"`
	Users       []User        `gorm:"many2many:master_class_users;"`
	Timelines   []Timeline    `gorm:"foreignKey:MasterClassID;references:ID"`
	Objectives  []Objective   `gorm:"foreignKey:MasterClassID;references:ID"`
	Presences   []Presence    `gorm:"foreignKey:MasterClassID;references:ID"`
	Tasks       []Task        `gorm:"foreignKey:MasterClassID;references:ID"`
	Materials   []Material    `gorm:"foreignKey:MasterClassID;references:ID"`
}
