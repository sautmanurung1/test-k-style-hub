package controller

import (
	"net/http"
	"skincare/interfaces"
	"skincare/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type productController struct {
	s interfaces.ProductService
}

func NewProductController(productService interfaces.ProductService) interfaces.ProductController {
	return &productController{
		s: productService,
	}
}

// InsertProduct implements interfaces.ProductController
func (h *productController) InsertProduct(c echo.Context) error {
	products := models.Product{
		ID: uuid.New().String(),
	}

	err := c.Bind(&products)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error to create product",
			"message": err.Error(),
		})
	}

	fails := h.s.InsertProduct(products)

	if fails != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error to insert data to database",
			"message": fails.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "success to insert product to database",
	})
}

// GetProduct implements interfaces.ProductController
func (h *productController) GetProduct(c echo.Context) error {
	id := c.Param("id")

	review, like, err := h.s.GetProduct(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error to get data from id",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success to get data from id",
		"review product": review,
		"counts likes":   like,
	})
}
