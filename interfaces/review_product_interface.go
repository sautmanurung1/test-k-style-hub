package interfaces

import (
	"skincare/models"

	"github.com/labstack/echo/v4"
)

type ReviewProductRepository interface {
	InsertReview(review models.ReviewProduct) error
}

type ReviewProductService interface {
	InsertReview(review models.ReviewProduct) error
}

type ReviewProductController interface {
	InsertReview(c echo.Context) error
}
