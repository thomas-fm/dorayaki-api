package database

import (
	// "encoding/csv"
	// "fmt"
	// "log"
	// "os"
	// "strconv"

	"dorayaki-api/models"

	"gorm.io/gorm"
)

const (
	VARIANTS_FILE string = "database/variants.csv"
	STORES_FILE   string = "database/stores.csv"
	STOCKS_FILE   string = "database/stocks.csv"
)

var (
	variants = []models.Variant{
		{
			Flavour:     "Nasi Padang",
			Description: "Dorayaki dengan cita rasa Indonesia",
			IMG_URL:     "https://res.cloudinary.com/stand-with-dorayaki/image/upload/v1627609112/download_1_jelwoz.jpg",
		},
		{
			Flavour:     "Ayam Geprek",
			Description: "Dorayaki dengan kombinasi ayam dan sambal geprek pedas kesukaan kita semua",
			IMG_URL:     "https://res.cloudinary.com/stand-with-dorayaki/image/upload/v1627609111/5f13dd959ae40_gj0c6m.jpg",
		},
		{
			Flavour:     "Nutella",
			Description: "Dorayaki yang dipadukan dengan nutella sebagai isian yang rasanya pasti maknyus",
			IMG_URL:     "https://res.cloudinary.com/stand-with-dorayaki/image/upload/v1627609110/images_rgvhmj.jpg",
		},
		{
			Flavour:     "Brown Sugar",
			Description: "Dorayaki dengan rasa brown sugar, memiliki kelezatan luar biasa",
			IMG_URL:     "https://res.cloudinary.com/stand-with-dorayaki/image/upload/v1627610897/dorayaki-hiroshi-yoshinaga-cc_txsgb5.jpg",
		},
	}
	stores = []models.Store{
		{
			Name:     "Dorayakee",
			Street:   "Jalan Mawar Nomor 10",
			District: "Coblong",
			Province: "Jawa Barat",
		},
		{
			Name:     "Dorayankee",
			Street:   "Jalan Melati Nomor 1",
			District: "Kiara Condong",
			Province: "Jawa Barat",
		},
		{
			Name:     "Doray",
			Street:   "Jalan Anggrek Nomor 13",
			District: "Coblong",
			Province: "Jawa Barat",
		},
		{
			Name:     "Doryaks",
			Street:   "Jalan Tulip Nomor 2",
			District: "Coblong",
			Province: "Jawa Barat",
		},
		{
			Name:     "Doryaks",
			Street:   "Jalan Matahari Nomor 3",
			District: "Antapani",
			Province: "Jawa Barat",
		},
	}
	stocks = []models.Stock{
		{
			StoreID:   1,
			VariantID: 1,
			Total:     100,
		},
		{
			StoreID:   1,
			VariantID: 2,
			Total:     300,
		},
		{
			StoreID:   1,
			VariantID: 3,
			Total:     600,
		},
		{
			StoreID:   2,
			VariantID: 1,
			Total:     150,
		},
		{
			StoreID:   3,
			VariantID: 1,
			Total:     250,
		},
		{
			StoreID:   4,
			VariantID: 3,
			Total:     356,
		},
		{
			StoreID:   5,
			VariantID: 2,
			Total:     124,
		},
	}
)

func Seed(db *gorm.DB) {
	// Insert to db
	var tokos []models.Store
	db.Find(&tokos)
	var dorayakis []models.Variant
	db.Find(&dorayakis)
	var quantity []models.Stock
	db.Find(&quantity)

	if len(tokos) == 0 && len(dorayakis) == 0 && len(quantity) == 0 {
		db.Create(&variants)
		db.Create(&stores)
		db.Create(&stocks)
	}
}
