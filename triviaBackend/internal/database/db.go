package database

import "gorm.io/gorm"

type Database interface {
	ConnectDB() *gorm.DB
}
