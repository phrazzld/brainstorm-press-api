package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}

	db.AutoMigrate(&Post{})

	return db
}
