package routes

import (
	"skincare/controller"
	"skincare/infrsatructure/database"
	"skincare/repository"
	"skincare/service"

	"github.com/labstack/echo/v4"
)

func RoutesProduct(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, conf)
	productController := controller.NewProductController(productService)

	product := echo.Group("/v1/product")
	product.POST("/create", productController.InsertProduct)
	product.GET("/:id", productController.GetProduct)
}
