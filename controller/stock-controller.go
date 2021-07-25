package controller

import (
	"net/http"
	"strconv"

	"dorayaki-api/dtos"
	"dorayaki-api/service"

	"dorayaki-api/models"
	"dorayaki-api/utils"

	"github.com/gin-gonic/gin"
)

type StockController interface {
	All(c *gin.Context)
	Insert(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindByVariantID(c *gin.Context)
}

type stockController struct {
	stockService service.StockService
}

func NewStockController(stockService service.StockService) StockController {
	return &stockController{
		stockService: stockService,
	}
}

// Get All by store id
func (s *stockController) All(c *gin.Context) {
	store_id, err := strconv.ParseUint(c.Param("storeID"), 0, 0)
	if err != nil {
		res := utils.BuildErrorResponse("No store id was found", err.Error(), utils.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		var stocks []models.Stock = s.stockService.Read(store_id)
		res := utils.BuildResponse(true, "OK", stocks)
		c.JSON(http.StatusOK, res)
	}
}

// Get by store id and variant id
func (s *stockController) FindByVariantID(c *gin.Context) {
	// Check if there is storeID
	store_id, err := strconv.ParseUint(c.Param("storeID"), 0, 0)
	if err != nil {
		res := utils.BuildErrorResponse("No store id was found", err.Error(), utils.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Check if there is variantID
	variant_id, err := strconv.ParseUint(c.Param("variantID"), 0, 0)
	if err != nil {
		res := utils.BuildErrorResponse("No variant id was found", err.Error(), utils.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Search what we are looking for in database
	var stock models.Stock = s.stockService.ReadByVId(store_id, variant_id)
	if (stock == models.Stock{}) {
		res := utils.BuildErrorResponse("No stock was found", err.Error(), utils.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		res := utils.BuildResponse(true, "OK", stock)
		c.JSON(http.StatusOK, res)
	}
}

func (s *stockController) Insert(c *gin.Context) {
	// var stockCreateDTO []dtos.StockSingleCreate
	var stockCreateDTO dtos.StockSingleCreate

	errDTO := c.ShouldBind(&stockCreateDTO)
	if errDTO != nil {
		res := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := s.stockService.Create(stockCreateDTO)
	res := utils.BuildResponse(true, "Success", result)
	c.JSON(http.StatusCreated, res)
}

func (s *stockController) Update(c *gin.Context) {
	var stockUpdateDTO dtos.StockSingleUpdate
	errDTO := c.ShouldBind(&stockUpdateDTO)
	if errDTO != nil {
		res := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := s.stockService.Update(stockUpdateDTO)
	res := utils.BuildResponse(true, "Success", result)
	c.JSON(http.StatusCreated, res)
}

func (s *stockController) Delete(c *gin.Context) {
	var stock models.Stock
	// Check if there is storeID
	store_id, err := strconv.ParseUint(c.Param("storeID"), 0, 0)
	if err != nil {
		res := utils.BuildErrorResponse("No store id was found", err.Error(), utils.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Check if there is variantID
	variant_id, err := strconv.ParseUint(c.Param("variantID"), 0, 0)
	if err != nil {
		res := utils.BuildErrorResponse("No variant id was found", err.Error(), utils.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	stock.StoreID = store_id
	stock.VariantID = variant_id

	s.stockService.Delete(stock)
	res := utils.BuildResponse(true, "Deleted", utils.EmptyObj{})
	c.JSON(http.StatusOK, res)
}
