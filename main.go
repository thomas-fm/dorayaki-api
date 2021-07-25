package main

import (
	"dorayaki-api/config"
	"dorayaki-api/controller"
	"dorayaki-api/database"
	"dorayaki-api/repository"
	"dorayaki-api/service"

	// "fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "net/http"
	"gorm.io/gorm"
)

var (
	// db
	db *gorm.DB = config.SetupDB()

	// repository
	variantRepository repository.VariantRepository = repository.NewVariantRepository(db)
	storeRepository   repository.StoreRepository   = repository.NewStoreRepository(db)
	stockRepository   repository.StockRepository   = repository.NewStockRepository(db)

	// service
	variantService service.VariantService = service.NewVariantService(variantRepository)
	storeService   service.StoreService   = service.NewStoreService(storeRepository)
	stockService   service.StockService   = service.NewStockService(stockRepository)

	// controller
	variantController controller.VariantController = controller.NewVariantController(variantService)
	storeController   controller.StoreController   = controller.NewStoreController(storeService)
	stockController   controller.StockController   = controller.NewStockController(stockService)
)

func main() {
	// Seeding
	database.Seed(db)
	r := gin.Default()
	r.Use(cors.Default())

	stockRoutes := r.Group("api/stocks")
	{
		stockRoutes.GET("/:storeID", stockController.All)
		stockRoutes.GET("/:storeID/variant/:variantID", stockController.FindByVariantID)
		// stockRoutes.POST("/:storeID/variant/:variantID")
		stockRoutes.POST("/:storeID", stockController.Insert)
		stockRoutes.PUT("/:storeID/variant/:variantID", stockController.Update)
		stockRoutes.DELETE("/:storeID/variant/:variantID", stockController.Delete)
	}

	storeRoutes := r.Group("api/stores")
	{
		storeRoutes.GET("/", storeController.All)
		storeRoutes.GET("/:id", storeController.FindByID)
		storeRoutes.POST("/", storeController.Insert)
		storeRoutes.PUT("/:id", storeController.Update)
		storeRoutes.DELETE("/:id", storeController.Delete)
	}
	variantRoutes := r.Group("api/variants")
	{
		variantRoutes.GET("/", variantController.All)
		variantRoutes.GET("/:id", variantController.FindByID)
		variantRoutes.POST("/", variantController.Insert)
		variantRoutes.PUT("/:id", variantController.Update)
		variantRoutes.DELETE("/:id", variantController.Delete)
	}

	defer config.CloseDatabaseConnection(db)
	r.Run()
}
