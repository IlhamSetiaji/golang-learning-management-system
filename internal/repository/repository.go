package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) FindAll(db *gorm.DB, entity *[]T) error {
	return db.Find(entity).Error
}

func (r *Repository[T]) FindByID(db *gorm.DB, entity *T, id int) error {
	return db.First(entity, id).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repository[T]) FindByField(db *gorm.DB, entity *[]T, field string, value string) error {
	return db.Where(field+" = ?", value).Find(entity).Error
}

func (r *Repository[T]) CountAll(db *gorm.DB, entity *[]T) int64 {
	var count int64
	db.Model(entity).Count(&count)
	return count
}

func (r *Repository[T]) CountByField(db *gorm.DB, entity *[]T, field string, value string) int64 {
	var count int64
	db.Model(entity).Where(field+" = ?", value).Count(&count)
	return count
}
