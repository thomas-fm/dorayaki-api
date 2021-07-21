package controller

import (
	"net/http"
	"strconv"

	"dorayaki-api/service"

	"dorayaki-api/dtos"
	"dorayaki-api/models"
	"dorayaki-api/utils"

	"github.com/gin-gonic/gin"
)

//BookController is a ...
type VariantController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindByID(context *gin.Context)
}

type variantController struct {
	variantService service.VariantService
}

//NewBookController create a new instances of BoookController
func NewVariantController(varService service.VariantService) VariantController {
	return &variantController{
		variantService: varService,
	}
}

// GET
func (v *variantController) All(context *gin.Context) {
	var variants []models.Variant = v.variantService.Read()
	res := utils.BuildResponse(true, "OK", variants)
	context.JSON(http.StatusOK, res)
}

// get by id
func (v *variantController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utils.BuildErrorResponse("No id was found", err.Error(), utils.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var variant models.Variant = v.variantService.ReadByID(id)
	if (variant == models.Variant{}) {
		res := utils.BuildErrorResponse("Data not found", "No data with given id", utils.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := utils.BuildResponse(true, "Success", variant)
		context.JSON(http.StatusOK, res)
	}
}

func (v *variantController) Insert(context *gin.Context) {
	var variantCreateDTO dtos.VariantCreateDTO
	errDTO := context.ShouldBind(&variantCreateDTO)
	if errDTO != nil {
		res := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := v.variantService.Create(variantCreateDTO)
		response := utils.BuildResponse(true, "Success", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (v *variantController) Update(context *gin.Context) {
	var variantUpdateDTO dtos.VariantUpdateDTO
	errDTO := context.ShouldBind(&variantUpdateDTO)
	if errDTO != nil {
		res := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result := v.variantService.Update(variantUpdateDTO)
	response := utils.BuildResponse(true, "Success", result)
	context.JSON(http.StatusOK, response)
}

func (v *variantController) Delete(context *gin.Context) {
	var variant models.Variant
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get id", "No param id were found", utils.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	variant.ID = id

	v.variantService.Delete(variant)
	res := utils.BuildResponse(true, "Deleted", utils.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
