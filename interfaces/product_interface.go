package interfaces

import (
	"skincare/models"

	"github.com/labstack/echo/v4"
)

type ProductRepository interface {
	InsertProduct(product models.Product) error
	GetProduct(id string) (review models.ReviewProduct, result int64, err error)
}

type ProductService interface {
	InsertProduct(product models.Product) error
	GetProduct(id string) (review models.ReviewProduct, result int64, err error)
}

type ProductController interface {
	InsertProduct(c echo.Context) error
	GetProduct(c echo.Context) error
}
