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
	Create(s dtos.StockCreateDTO) models.Stock
	Update(s dtos.StockUpdateDTO) models.Stock
	Delete(s models.Stock)
	Read() []models.Stock
	ReadByID(IDStock uint64) models.Stock
}

type stockService struct {
	stockRepository repository.StockRepository
}

// //NewBookService .....
// func NewStockService(stockRepo repository.StockRepository) StockService {
// 	return &stockService{
// 		stockRepository: stockRepo,
// 	}
// }

func (service *stockService) Create(s dtos.StockSingleCreate) models.Stock {
	stock := models.Stock{}
	err := copier.Copy(&stock, &s)
	if err != nil {
		log.Fatalf("Failed copy %v: ", err)
		// os.Exit(1)
	}
	res := service.stockRepository.CreateStock(stock)
	return res
}

func (service *stockService) Update(s dtos.StockUpdateDTO) models.Stock {
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

func (service *stockService) Read() []models.Stock {
	return service.stockRepository.ReadStocks()
}

func (service *stockService) ReadByID(IDStock uint64) models.Stock {
	return service.stockRepository.ReadStockByID(IDStock)
}

func (service *stockService) Filter(s dtos.StockFilterDTO) []models.Stock {
	return service.stockRepository.Filter(s.District, s.Province)
}
