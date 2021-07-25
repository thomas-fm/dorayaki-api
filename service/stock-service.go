package service

import (
	"log"

	"dorayaki-api/dtos"
	"dorayaki-api/models"
	"dorayaki-api/repository"

	"github.com/jinzhu/copier"
)

//BookService is a ....
type StockService interface {
	// Create(s []dtos.StockSingleCreate) []models.Stock
	Create(s dtos.StockSingleCreate) models.Stock

	Update(s dtos.StockSingleUpdate) models.Stock
	Delete(s models.Stock)
	Read(store_id uint64) []models.Stock
	ReadByVId(store_id uint64, variant_id uint64) models.Stock
}

type stockService struct {
	stockRepository repository.StockRepository
}

func NewStockService(stockRepo repository.StockRepository) StockService {
	return &stockService{
		stockRepository: stockRepo,
	}
}

// func (service *stockService) Create(stocksDTO []dtos.StockSingleCreate) []models.Stock {
// 	var stocks []models.Stock
// 	err := copier.Copy(&stocks, &stocksDTO)
// 	if err != nil {
// 		log.Fatalf("Failed copy %v: ", err)
// 	}
// 	for _, stock := range stocks {
// 		service.stockRepository.CreateStock(stock)
// 	}
// 	return stocks
// }

func (service *stockService) Create(stocksDTO dtos.StockSingleCreate) models.Stock {
	var stocks models.Stock
	err := copier.Copy(&stocks, &stocksDTO)
	if err != nil {
		log.Fatalf("Failed copy %v: ", err)
	}
	res := service.stockRepository.CreateStock(stocks)

	return res
}

func (service *stockService) Update(s dtos.StockSingleUpdate) models.Stock {
	stock := models.Stock{}
	err := copier.Copy(&stock, &s)
	if err != nil {
		log.Fatalf("Failed copy %v: ", err)
	}
	res := service.stockRepository.UpdateStock(stock)
	return res
}

func (service *stockService) Delete(s models.Stock) {
	service.stockRepository.DeleteStock(s)
}

func (service *stockService) Read(store_id uint64) []models.Stock {
	return service.stockRepository.ReadStock(store_id)
}

func (service *stockService) ReadByVId(store_id uint64, variant_id uint64) models.Stock {
	return service.stockRepository.ReadStockByVid(store_id, variant_id)
}
