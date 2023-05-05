package routes

import (
	"skincare/controller"
	"skincare/infrsatructure/database"
	"skincare/repository"
	"skincare/service"

	"github.com/labstack/echo/v4"
)

func ReviewProductRoutes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	reviewProductRepository := repository.NewReviewProductRepository(db)
	reviewProductService := service.NewReviewProductService(reviewProductRepository, conf)
	reviewProductController := controller.NewReviewProductController(reviewProductService)

	reviewProduct := echo.Group("/v1/review")
	reviewProduct.POST("/create", reviewProductController.InsertReview)
}
