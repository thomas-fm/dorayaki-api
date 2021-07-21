package service

import (
	"log"

	"dorayaki-api/dtos"
	"dorayaki-api/models"
	"dorayaki-api/repository"

	"github.com/jinzhu/copier"
)

//BookService is a ....
type StoreService interface {
	Create(s dtos.StoreCreateDTO) models.Store
	Update(s dtos.StoreUpdateDTO) models.Store
	Delete(s models.Store)
	Read() []models.Store
	ReadByID(IDStore uint64) models.Store
	Filter(s dtos.StoreFilterDTO) []models.Store
}

type storeService struct {
	storeRepository repository.StoreRepository
}

//NewBookService .....
func NewStoreService(storeRepo repository.StoreRepository) StoreService {
	return &storeService{
		storeRepository: storeRepo,
	}
}

func (service *storeService) Create(s dtos.StoreCreateDTO) models.Store {
	store := models.Store{}
	err := copier.Copy(&store, &s)
	if err != nil {
		log.Fatalf("Failed copy %v: ", err)
		// os.Exit(1)
	}
	res := service.storeRepository.CreateStore(store)
	return res
}

func (service *storeService) Update(s dtos.StoreUpdateDTO) models.Store {
	store := models.Store{}
	err := copier.Copy(&store, &s)
	if err != nil {
		log.Fatalf("Failed copy %v: ", err)
	}
	res := service.storeRepository.UpdateStore(store)
	return res
}

func (service *storeService) Delete(s models.Store) {
	service.storeRepository.DeleteStore(s)
}

func (service *storeService) Read() []models.Store {
	return service.storeRepository.ReadStores()
}

func (service *storeService) ReadByID(IDStore uint64) models.Store {
	return service.storeRepository.ReadStoreByID(IDStore)
}

func (service *storeService) Filter(s dtos.StoreFilterDTO) []models.Store {
	return service.storeRepository.Filter(s.District, s.Province)
}
