package repository

import (
	"errors"
	"time"

	"github.com/IlhamSetiaji/go-lms/internal/entity"
	"github.com/IlhamSetiaji/go-lms/internal/request"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

type UserRepositoryInterface interface {
	FindByField(entity *[]entity.User, field string, value string) (*[]entity.User, error)
	FindFirstByField(db *gorm.DB, entity *entity.User, field string, value string) (*entity.User, error)
	CountAll(entity *entity.User) (int64, error)
	CountByField(entity *entity.User, field string, value string) (int64, error)
	GetMyProfile(db *gorm.DB, entity *entity.User, id uint) (*entity.User, error)
	CreateUser(db *gorm.DB, entity *entity.User, payload *request.RegisterUserRequest) (*entity.User, error)
	AssignRole(db *gorm.DB, userId uint, roleId uint) error
}

func NewUserRepository(db *gorm.DB, log *logrus.Logger) UserRepositoryInterface {
	return &UserRepository{
		DB:  db,
		Log: log,
	}
}

func (r *UserRepository) FindByField(entity *[]entity.User, field string, value string) (*[]entity.User, error) {
	return entity, r.DB.Where(field+" = ?", value).Find(entity).Error
}

func (r *UserRepository) FindFirstByField(db *gorm.DB, entity *entity.User, field string, value string) (*entity.User, error) {
	if db == nil {
		return nil, errors.New("database connection is not initialized")
	}

	err := db.Where(field+" = ?", value).First(entity).Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *UserRepository) CountAll(entity *entity.User) (int64, error) {
	var count int64
	err := r.DB.Model(entity).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountByField(entity *entity.User, field string, value string) (int64, error) {
	var count int64
	err := r.DB.Model(entity).Where(field+" = ?", value).Count(&count).Error
	return count, err
}

func (r *UserRepository) GetMyProfile(db *gorm.DB, entity *entity.User, id uint) (*entity.User, error) {
	return entity, db.Preload("Roles").First(entity, "users.id = ?", id).Error
}

func (r *UserRepository) CreateUser(db *gorm.DB, entity *entity.User, payload *request.RegisterUserRequest) (*entity.User, error) {
	entity.Name = payload.Name
	entity.Username = payload.Username
	entity.Email = payload.Email
	entity.Password = payload.Password

	result := db.Create(entity)
	return entity, result.Error
}

func (r *UserRepository) AssignRole(db *gorm.DB, userId uint, roleId uint) error {
	err := db.Model(&entity.User{ID: userId}).Association("Roles").Append(&entity.Role{Model: gorm.Model{ID: roleId, CreatedAt: time.Now(), UpdatedAt: time.Now()}})
	return err
}
