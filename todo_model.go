package main

import "github.com/jinzhu/gorm"

type (
	// todoModel describes a todoModel type
	todoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}
)

func migrateModel(db *gorm.DB) {
	Log.Printf("Migrating todo model")
	db.AutoMigrate(&todoModel{})
}
