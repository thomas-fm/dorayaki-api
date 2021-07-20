package repository

import (
	"github.com/thomas-fm/dorayaki-api/models"

	"gorm.io/gorm"
)

type StoreRepository interface {
	CreateStore(s models.Store) models.Store
	UpdateStore(s models.Store) models.Store
	DeleteStore(s models.Store)
	ReadStores() []models.Store
	ReadStoreByID(IDStore uint64) models.Store
	Filter(district string, province string) []models.Store
}

type storeConnection struct {
	connection *gorm.DB
}

func NewStoreRepository(dbConn *gorm.DB) StoreRepository {
	return &storeConnection{
		connection: dbConn,
	}
}

// CREATE
func (db *storeConnection) CreateStore(s models.Store) models.Store {
	db.connection.Save(&s)
	db.connection.Find(&s)
	return s
}

// READ
func (db *storeConnection) ReadStores() []models.Store {
	var stores []models.Store
	db.connection.Find(&stores)

	return stores
}

// UPDATE
func (db *storeConnection) UpdateStore(s models.Store) models.Store {
	db.connection.Save(&s)
	db.connection.Find(&s)
	return s
}

// DELETE
func (db *storeConnection) DeleteStore(s models.Store) {
	db.connection.Delete(&s)
}

// find by id
func (db *storeConnection) ReadStoreByID(IDStore uint64) models.Store {
	var store models.Store
	db.connection.Find(&store, IDStore)
	return store
}
func (db *storeConnection) Filter(district string, province string) (stores []models.Store) {
	db.connection.Where("district = ? AND province = ?", district, province).Find(&stores)
	return stores
}
