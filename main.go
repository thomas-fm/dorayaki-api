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

	// service
	variantService service.VariantService = service.NewVariantService(variantRepository)

	// controller
	variantController controller.VariantController = controller.NewVariantController(variantService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	variantRoutes := r.Group("api/variants")
	{
		variantRoutes.GET("/", variantController.All)
		variantRoutes.GET("/:id", variantController.FindByID)
		variantRoutes.POST("/", variantController.Insert)
		variantRoutes.PUT("/:id", variantController.Update)
		variantRoutes.DELETE("/", variantController.Delete)
	}

	r.Run()
}
