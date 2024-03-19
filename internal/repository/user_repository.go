package repository

import (
	"github.com/IlhamSetiaji/go-lms/internal/entity"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

type UserRepositoryInterface interface {
	Create(entity *entity.User) (*entity.User, error)
	FindAll(entity *[]entity.User) (*[]entity.User, error)
	FindByID(entity *entity.User, id int) (*entity.User, error)
	Update(entity *entity.User) (*entity.User, error)
	Delete(entity *entity.User) (*entity.User, error)
	FindByField(entity *[]entity.User, field string, value string) (*[]entity.User, error)
	FindFirstByField(entity *entity.User, field string, value string) (*entity.User, error)
	CountAll(entity *entity.User) (int64, error)
	CountByField(entity *entity.User, field string, value string) (int64, error)
}

func NewUserRepository(log *logrus.Logger) UserRepositoryInterface {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) Create(entity *entity.User) (*entity.User, error) {
	return entity, r.DB.Create(entity).Error
}

func (r *UserRepository) FindAll(entity *[]entity.User) (*[]entity.User, error) {
	return entity, r.DB.Find(entity).Error
}

func (r *UserRepository) FindByID(entity *entity.User, id int) (*entity.User, error) {
	return entity, r.DB.First(entity, id).Error
}

func (r *UserRepository) Update(entity *entity.User) (*entity.User, error) {
	return entity, r.DB.Save(entity).Error
}

func (r *UserRepository) Delete(entity *entity.User) (*entity.User, error) {
	return entity, r.DB.Delete(entity).Error
}

func (r *UserRepository) FindByField(entity *[]entity.User, field string, value string) (*[]entity.User, error) {
	return entity, r.DB.Where(field+" = ?", value).Find(entity).Error
}

func (r *UserRepository) FindFirstByField(entity *entity.User, field string, value string) (*entity.User, error) {
	return entity, r.DB.Where(field+" = ?", value).First(entity).Error
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
