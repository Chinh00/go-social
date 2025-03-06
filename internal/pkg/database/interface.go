package database

import "gorm.io/gorm"

type DataContext interface {
	GetDB() *gorm.DB
	Configure(...Options)
	Migrations(models ...interface{}) error
}
