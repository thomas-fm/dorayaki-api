package main

import (
	"dorayaki-api/models"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	// Drop table
	db.Migrator().DropTable(&models.Variant{}, &models.Store{}, &models.Stock{})

	// Migrate the type
	db.Debug().AutoMigrate(&models.Variant{}, &models.Store{}, &models.Stock{})
}
