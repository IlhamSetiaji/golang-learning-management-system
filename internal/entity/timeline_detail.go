package entity

import "gorm.io/gorm"

type TimelineDetail struct {
	gorm.Model
	TimelineID uint     `json:"timeline_id"`
	Timeline   Timeline `gorm:"foreignKey:TimelineID;references:ID"`
	Activity   string   `json:"activity"`
}
