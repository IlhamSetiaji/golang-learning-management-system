package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}
