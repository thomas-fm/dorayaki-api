package repository

import (
	"github.com/thomas-fm/dorayaki-api/models"

	"gorm.io/gorm"
)

type VariantRepository interface {
	InsertVariant(v models.Variant) models.Variant
	UpdateVariant(v models.Variant) models.Variant
	DeleteVariant(v models.Variant)
	AllVariants() []models.Variant
	FindVariantbyID(IDVariant uint64) models.Variant
}

type variantConnection struct {
	connection *gorm.DB
}

func NewVariantRepository(dbConn *gorm.DB) VariantRepository {
	return &variantConnection{
		connection: dbConn,
	}
}

// CREATE
func (db *variantConnection) InsertVariant(v models.Variant) models.Variant {
	db.connection.Save(&v)
	db.connection.Find(&v)
	return v
}

// READ
func (db *variantConnection) AllVariants() []models.Variant {
	var variants []models.Variant
	db.connection.Find(&variants)

	return variants
}

// UPDATE
func (db *variantConnection) UpdateVariant(v models.Variant) models.Variant {
	db.connection.Save(&v)
	db.connection.Find(&v)
	return v
}

// DELETE
func (db *variantConnection) DeleteVariant(v models.Variant) {
	db.connection.Delete(&v)
}

// find by id
func (db *variantConnection) FindVariantbyID(IDVariant uint64) models.Variant {
	var variant models.Variant
	db.connection.Find(&variant, IDVariant)
	return variant
}
