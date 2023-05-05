package routes

import (
	"skincare/controller"
	"skincare/infrsatructure/database"
	"skincare/repository"
	"skincare/service"

	"github.com/labstack/echo/v4"
)

func LikeReviewRoutes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	likeReviewRepository := repository.NewLikeReviewRepository(db)
	likeReviewService := service.NewLikeReviewService(likeReviewRepository, conf)
	likeReviewController := controller.NewLikeReviewController(likeReviewService)

	likeReview := echo.Group("/v1/like")
	likeReview.POST("/create", likeReviewController.InsertLikeReview)
	likeReview.DELETE("/dislike/:id", likeReviewController.DeleteLikeReview)
}
