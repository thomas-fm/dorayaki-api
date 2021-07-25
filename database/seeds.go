package database

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"dorayaki-api/models"

	"gorm.io/gorm"
)

const (
	VARIANTS_FILE string = "database/variants.csv"
	STORES_FILE   string = "database/stores.csv"
	STOCKS_FILE   string = "database/stocks.csv"
)

var (
	variants []models.Variant
	stores   []models.Store
	stocks   []models.Stock
)

func Seed(db *gorm.DB) {
	variantSeed, err1 := readCSV(VARIANTS_FILE)
	storeSeed, err2 := readCSV(STORES_FILE)
	stockSeed, err3 := readCSV(STOCKS_FILE)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1.Error())
		fmt.Println(err2.Error())
		fmt.Println(err3.Error())
		log.Fatal("couldnt read file")
	}

	// Insert into batch
	// Variants seed
	for _, line := range variantSeed {
		data := models.Variant{
			Flavour:     line[0],
			Description: line[1],
			IMG_URL:     line[2],
		}

		variants = append(variants, data)
	}
	// Store seed
	for _, line := range storeSeed {
		data := models.Store{
			Name:     line[0],
			Street:   line[1],
			District: line[2],
			Province: line[3],
			// Phone_Number: line[4],
		}

		stores = append(stores, data)
	}
	// Stock seed
	for _, line := range stockSeed {
		s, err := strconv.ParseUint(line[0], 10, 64)
		v, err := strconv.ParseUint(line[1], 10, 64)
		t, err := strconv.ParseUint(line[2], 10, 64)

		if err != nil {
			log.Fatal("cannot read file")
		}
		data := models.Stock{
			StoreID:   s,
			VariantID: v,
			Total:     t,
		}
		stocks = append(stocks, data)
	}

	// Insert to db
	db.Create(&variants)
	db.Create(&stores)
	db.Create(&stocks)
}

// read csv file return as 2d array
func readCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
