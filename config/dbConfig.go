package config

import (
	"fmt"
	"log"
	"os"

	"dorayaki-api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// const (
// 	dbUser = "root"
// 	dbPass = ""
// 	dbHost = "localhost"
// 	dbName = "dorayaki_api"
// )

var (
	dbUser string
	dbPass string
	dbHost string
	dbName string
)

func SetupDB() (db *gorm.DB) {
	var err error
	err = godotenv.Load()

	dbUser = os.Getenv("DBUser")
	dbPass = os.Getenv("DBPass")
	dbHost = os.Getenv("DBHost")
	dbName = os.Getenv("DBName")

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
