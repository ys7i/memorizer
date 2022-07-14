package usecase

import "gorm.io/gorm"

type DBRepository interface {
	Begin() *gorm.DB
	Connect() *gorm.DB
}