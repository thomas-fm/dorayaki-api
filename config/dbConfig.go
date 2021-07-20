package config

import (
	"fmt"

	"github.com/thomas-fm/dorayaki-api/models"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUser = "root"
	dbPass = ""
	dbHost = "localhost"
	dbName = "dorayaki_api"
)

func SetupDB() (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	// migrate models to database
	db.Debug().AutoMigrate(&models.Variant{}, &models.Store{}, &models.Stock{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
