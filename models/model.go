package models

import (
	"gorm.io/gorm"
)

// Table of dorayaki stock from each store
type Stock struct {
	gorm.Model
	StoreID   uint64  `gorm:"primary_key"`
	VariantID uint64  `gorm:"primary_key"`
	Store     Store   `gorm:"foreignKey:StoreID"`
	Variant   Variant `gorm:"foreignKey:VariantID"`
	Total     uint64
}

// Dorayaki variant table
type Variant struct {
	gorm.Model
	ID          uint64 `gorm:"primary_key;autoIncrement:true" json:"id"`
	Flavour     string `gorm:"type:varchar(255)" json:"flavour"`
	Description string `gorm:"type:text" json:"description"`
	IMG_URL     string `gorm:"type:text" json:"img_url"`
}

// Table of dorayaki store
type Store struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name         string
	Street       string
	District     string
	Province     string
	Phone_Number string
}

// Table of user
type User struct {
	ID     int64 `gorm:"primaryKey"`
	Name   string
	Status string // admin or superadmin
}