package main

import (
	"github.com/thomas-fm/dorayaki-api/config"
	"github.com/thomas-fm/dorayaki-api/controller"
	"github.com/thomas-fm/dorayaki-api/repository"
	"github.com/thomas-fm/dorayaki-api/service"

	// "fmt"

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

	// service
	variantService service.VariantService = service.NewVariantService(variantRepository)
	storeService   service.StoreService   = service.NewStoreService(storeRepository)

	// controller
	variantController controller.VariantController = controller.NewVariantController(variantService)
	storeController   controller.StoreController   = controller.NewStoreController(storeService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

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

	r.Run()
}
