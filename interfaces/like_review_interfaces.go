package interfaces

import (
	"skincare/models"

	"github.com/labstack/echo/v4"
)

type LikeReviewRepository interface {
	InsertLikeReview(like models.LikeReview) error
	DeleteLikeReview(id string, like models.LikeReview)
}

type LikeReviewService interface {
	InsertLikeReview(like models.LikeReview) error
	DeleteLikeReview(id string, like models.LikeReview)
}

type LikeReviewController interface {
	InsertLikeReview(c echo.Context) error
	DeleteLikeReview(c echo.Context) error
}
