package service

import (
	"log"

	"github.com/thomas-fm/dorayaki-api/dtos"
	"github.com/thomas-fm/dorayaki-api/models"
	"github.com/thomas-fm/dorayaki-api/repository"

	"github.com/mashingan/smapping"
)

//BookService is a ....
type VariantService interface {
	Insert(v dtos.VariantCreateDTO) models.Variant
	Update(v dtos.VariantUpdateDTO) models.Variant
	Delete(v models.Variant)
	All() []models.Variant
	FindByID(IDVariant uint64) models.Variant
	// IsAllowedToEdit(userID string, bookID uint64) bool
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

func (service *variantService) Insert(v dtos.VariantCreateDTO) models.Variant {
	variant := models.Variant{}
	err := smapping.FillStruct(&variant, smapping.MapFields(&v))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.variantRepository.InsertVariant(variant)
	return res
}

func (service *variantService) Update(v dtos.VariantUpdateDTO) models.Variant {
	variant := models.Variant{}
	err := smapping.FillStruct(&variant, smapping.MapFields(&v))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.variantRepository.UpdateVariant(variant)
	return res
}

func (service *variantService) Delete(v models.Variant) {
	service.variantRepository.DeleteVariant(v)
}

func (service *variantService) All() []models.Variant {
	return service.variantRepository.AllVariants()
}

func (service *variantService) FindByID(IDVariant uint64) models.Variant {
	return service.variantRepository.FindVariantbyID(IDVariant)
}
