package main

import (
	"github.com/IlhamSetiaji/go-lms/internal/config"
	"github.com/IlhamSetiaji/go-lms/internal/entity"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logger
	log := logrus.New()

	// Initialize Viper
	v := config.NewViper()
	// Initialize database
	db := config.NewDatabase(v, log)

	// Migrate the schema
	err := db.AutoMigrate(&entity.User{}, &entity.PasswordResetToken{}, &entity.Role{}, &entity.UserRole{}, &entity.ProgramCategory{}, &entity.Program{},
		&entity.Batch{}, &entity.MasterClass{}, &entity.MasterClassUser{}, &entity.Objective{}, &entity.Presence{}, &entity.Timeline{}, &entity.Material{}, &entity.Task{},
		&entity.TimelineDetail{}, &entity.PresenceUser{}, &entity.TaskAttachment{}, &entity.MenteeTask{}, &entity.MenteeTaskAttachment{}, &entity.MaterialAttachment{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
