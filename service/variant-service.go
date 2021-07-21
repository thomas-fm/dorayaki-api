package service

import (
	"log"

	"dorayaki-api/dtos"
	"dorayaki-api/models"
	"dorayaki-api/repository"

	"github.com/jinzhu/copier"
)

//BookService is a ....
type VariantService interface {
	Create(v dtos.VariantCreateDTO) models.Variant
	Update(v dtos.VariantUpdateDTO) models.Variant
	Delete(v models.Variant)
	Read() []models.Variant
	ReadByID(IDVariant uint64) models.Variant
	// IsReadowedToEdit(userID string, bookID uint64) bool
}

type variantService struct {
	variantRepository repository.VariantRepository
}

//NewBookService .....
func NewVariantService(varRepo repository.VariantRepository) VariantService {
	return &variantService{
		variantRepository: varRepo,
	}
}

func (service *variantService) Create(v dtos.VariantCreateDTO) models.Variant {
	variant := models.Variant{}
	err := copier.Copy(&variant, &v)
	if err != nil {
		log.Fatalf("Failed copy %v: ", err)
	}
	res := service.variantRepository.CreateVariant(variant)
	return res
}

func (service *variantService) Update(v dtos.VariantUpdateDTO) models.Variant {
	variant := models.Variant{}
	err := copier.Copy(&variant, &v)
	if err != nil {
		log.Fatalf("Failed copy %v: ", err)
	}
	res := service.variantRepository.UpdateVariant(variant)
	return res
}

func (service *variantService) Delete(v models.Variant) {
	service.variantRepository.DeleteVariant(v)
}

func (service *variantService) Read() []models.Variant {
	return service.variantRepository.ReadVariants()
}

func (service *variantService) ReadByID(IDVariant uint64) models.Variant {
	return service.variantRepository.ReadVariantbyID(IDVariant)
}
