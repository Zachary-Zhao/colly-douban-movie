package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name     string
	Director string
	Year     int8
	Score    float32
}

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Movie{})
}

func BulkCreate(movies []*Movie) (b bool, num int) {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.CreateInBatches(movies, 100)
	return
}
