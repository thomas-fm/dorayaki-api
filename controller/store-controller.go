package controller

import (
	"net/http"
	"strconv"

	"github.com/thomas-fm/dorayaki-api/service"

	"github.com/thomas-fm/dorayaki-api/dtos"
	"github.com/thomas-fm/dorayaki-api/models"
	"github.com/thomas-fm/dorayaki-api/utils"

	"github.com/gin-gonic/gin"
)

//BookController is a ...
type StoreController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindByID(context *gin.Context)
}

type storeController struct {
	storeService service.StoreService
}

//NewBookController create a new instances of BoookController
func NewStoreController(varService service.StoreService) StoreController {
	return &storeController{
		storeService: varService,
	}
}

// GET
func (s *storeController) All(context *gin.Context) {
	// handle get with filter
	var storeFilterDTO dtos.StoreFilterDTO
	errDTO := context.ShouldBind(&storeFilterDTO)

	if errDTO != nil {
		var stores []models.Store = s.storeService.Read()
		res := utils.BuildResponse(true, "OK", stores)
		context.JSON(http.StatusOK, res)
	} else {
		var stores []models.Store = s.storeService.Filter(storeFilterDTO)
		res := utils.BuildResponse(true, "OK", stores)
		context.JSON(http.StatusOK, res)
	}

}

// get by id
func (s *storeController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utils.BuildErrorResponse("No id was found", err.Error(), utils.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var store models.Store = s.storeService.ReadByID(id)
	if (store == models.Store{}) {
		res := utils.BuildErrorResponse("Data not found", "No data with given id", utils.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := utils.BuildResponse(true, "Success", store)
		context.JSON(http.StatusOK, res)
	}
}

func (s *storeController) Insert(context *gin.Context) {
	var storeCreateDTO dtos.StoreCreateDTO
	errDTO := context.ShouldBind(&storeCreateDTO)
	if errDTO != nil {
		res := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := s.storeService.Create(storeCreateDTO)
		response := utils.BuildResponse(true, "Success", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (s *storeController) Update(context *gin.Context) {
	var storeUpdateDTO dtos.StoreUpdateDTO
	errDTO := context.ShouldBind(&storeUpdateDTO)
	if errDTO != nil {
		res := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result := s.storeService.Update(storeUpdateDTO)
	response := utils.BuildResponse(true, "Success", result)
	context.JSON(http.StatusOK, response)
}

func (s *storeController) Delete(context *gin.Context) {
	var store models.Store
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get id", "No param id were found", utils.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	store.ID = id

	s.storeService.Delete(store)
	res := utils.BuildResponse(true, "Deleted", utils.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
