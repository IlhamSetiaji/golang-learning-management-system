package main

import (
	"time"

	"github.com/IlhamSetiaji/go-lms/internal/config"
	"github.com/IlhamSetiaji/go-lms/internal/entity"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

	// Seed the database

	// Create roles
	roles := []*entity.Role{
		{Name: "admin"},
		{Name: "mentor"},
		{Name: "mentee"},
	}
	db.Create(roles)

	// Create users
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte("changeme"), bcrypt.DefaultCost)
	users := []*entity.User{
		{Email: "admin@test.test", Password: string(hashedPasswordBytes), Name: "Admin", Username: "admin", IsActive: true, EmailVerifiedAt: time.Now()},
		{Email: "mentor@test.test", Password: string(hashedPasswordBytes), Name: "Mentor", Username: "mentor", IsActive: true, EmailVerifiedAt: time.Now()},
		{Email: "mentee@test.test", Password: string(hashedPasswordBytes), Name: "Mentee", Username: "mentee", IsActive: true, EmailVerifiedAt: time.Now()},
	}
	db.Create(users)

	// Assign roles to users
	userRoles := []*entity.UserRole{
		{UserID: 1, RoleID: 1},
		{UserID: 2, RoleID: 2},
		{UserID: 3, RoleID: 3},
	}
	db.Create(userRoles)

	// Create program categories
	programCategories := []*entity.ProgramCategory{
		{Name: "Bootcamp", Description: "Bootcamp description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
		{Name: "Corporate Training", Description: "Corporate Training description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
		{Name: "Workshop", Description: "Workshop description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
		{Name: "Webinar", Description: "Webinar description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
		{Name: "Crash Course", Description: "Crash Course description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
	}
	db.Create(programCategories)

	// Create programs
	programs := []*entity.Program{
		{ProgramCategoryID: 1, Name: "Magang Online Academy", Description: "Magang Online Academy description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
		{ProgramCategoryID: 2, Name: "Corporate Training", Description: "Corporate Training description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
		{ProgramCategoryID: 3, Name: "Workshop", Description: "Workshop description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
		{ProgramCategoryID: 4, Name: "Webinar", Description: "Webinar description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
		{ProgramCategoryID: 5, Name: "Crash Course", Description: "Crash Course description", Image: "https://source.unsplash.com/random/900x700/?fruit"},
	}
	db.Create(programs)

	// Create batches
	batches := []*entity.Batch{
		{ProgramID: 1, Name: "Magang Online Academy", Description: "Batch 1", Image: "https://source.unsplash.com/random/900x700/?fruit", RegistrationStartDate: time.Now(), RegistrationEndDate: time.Now().AddDate(0, 0, 7), StartDate: time.Now().AddDate(0, 0, 14), EndDate: time.Now().AddDate(0, 0, 28)},
		{ProgramID: 2, Name: "Corporate Training", Description: "Batch 1", Image: "https://source.unsplash.com/random/900x700/?fruit", RegistrationStartDate: time.Now(), RegistrationEndDate: time.Now().AddDate(0, 0, 7), StartDate: time.Now().AddDate(0, 0, 14), EndDate: time.Now().AddDate(0, 0, 28)},
		{ProgramID: 3, Name: "Workshop", Description: "Batch 1", Image: "https://source.unsplash.com/random/900x700/?fruit", RegistrationStartDate: time.Now(), RegistrationEndDate: time.Now().AddDate(0, 0, 7), StartDate: time.Now().AddDate(0, 0, 14), EndDate: time.Now().AddDate(0, 0, 28)},
		{ProgramID: 4, Name: "Webinar", Description: "Batch 1", Image: "https://source.unsplash.com/random/900x700/?fruit", RegistrationStartDate: time.Now(), RegistrationEndDate: time.Now().AddDate(0, 0, 7), StartDate: time.Now().AddDate(0, 0, 14), EndDate: time.Now().AddDate(0, 0, 28)},
		{ProgramID: 5, Name: "Crash Course", Description: "Batch 1", Image: "https://source.unsplash.com/random/900x700/?fruit", RegistrationStartDate: time.Now(), RegistrationEndDate: time.Now().AddDate(0, 0, 7), StartDate: time.Now().AddDate(0, 0, 14), EndDate: time.Now().AddDate(0, 0, 28)},
	}
	db.Create(batches)

	// Create master classes
	masterClasses := []*entity.MasterClass{
		{BatchID: 1, Name: "Web Development", Slug: "web-development", Description: "Web Development description", Image: "https://source.unsplash.com/random/900x700/?fruit", Price: 100000, Duration: 60, Level: "Beginner", IsActive: true},
		{BatchID: 2, Name: "Android Development", Slug: "android-development", Description: "Android Development description", Image: "https://source.unsplash.com/random/900x700/?fruit", Price: 100000, Duration: 60, Level: "Beginner", IsActive: true},
		{BatchID: 3, Name: "Flutter Development", Slug: "flutter-development", Description: "Flutter Development description", Image: "https://source.unsplash.com/random/900x700/?fruit", Price: 100000, Duration: 60, Level: "Beginner", IsActive: true},
		{BatchID: 4, Name: "Business Analyst", Slug: "business-analyst", Description: "Business Analyst description", Image: "https://source.unsplash.com/random/900x700/?fruit", Price: 100000, Duration: 60, Level: "Beginner", IsActive: true},
		{BatchID: 5, Name: "UI/UX Design", Slug: "ui-ux-design", Description: "UI/UX Design description", Image: "https://source.unsplash.com/random/900x700/?fruit", Price: 100000, Duration: 60, Level: "Beginner", IsActive: true},
	}
	db.Create(masterClasses)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
