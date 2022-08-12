package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectionDb() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	return db, err
}
