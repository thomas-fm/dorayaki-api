package repository

import (
	"dorayaki-api/models"

	"gorm.io/gorm"
)

type StockRepository interface {
	DeleteStock(s models.Stock)
	UpdateStock(s models.Stock) models.Stock
	CreateStock(s models.Stock) models.Stock
	ReadStock(store_id uint64) (stocks []models.Stock)
	ReadStockByVid(store_id uint64, variant_id uint64) (stock models.Stock)
}

type stockConnection struct {
	connection *gorm.DB
}

func NewStockRepository(dbConn *gorm.DB) StockRepository {
	return &stockConnection{
		connection: dbConn,
	}
}
func (db *stockConnection) ReadStock(store_id uint64) (stocks []models.Stock) {
	db.connection.Where("store_id = ?", store_id).Find(&stocks)
	return stocks
}

func (db *stockConnection) ReadStockByVid(store_id uint64, variant_id uint64) (stock models.Stock) {
	db.connection.Where("store_id = ? AND variant_id = ?", store_id, variant_id).Find(&stock)
	return stock
}

func (db *stockConnection) CreateStock(s models.Stock) models.Stock {
	db.connection.Save(&s)
	db.connection.Find(&s)

	return s
}

func (db *stockConnection) UpdateStock(s models.Stock) models.Stock {
	db.connection.Save(&s)
	db.connection.Find(&s)
	return s
}

func (db *stockConnection) DeleteStock(s models.Stock) {
	db.connection.Delete(&s)
}

// // CREATE
// func (db *stockConnection) CreateStock(v models.Stock) models.Stock {
// 	db.connection.Save(&v)
// 	db.connection.Find(&v)
// 	return v
// }

// // READ
// func (db *stockConnection) ReadStocks() []models.Stock {
// 	var stocks []models.Stock
// 	db.connection.Find(&stocks)

// 	return stocks
// }

// // UPDATE
// func (db *stockConnection) UpdateStock(v models.Stock) models.Stock {
// 	db.connection.Save(&v)
// 	db.connection.Find(&v)
// 	return v
// }

// // DELETE
// func (db *stockConnection) DeleteStock(v models.Stock) {
// 	db.connection.Delete(&v)
// }

// // find by id
// func (db *stockConnection) ReadStockbyID(IDStock uint64) models.Stock {
// 	var stock models.Stock
// 	db.connection.Find(&stock, IDStock)
// 	return stock
// }
