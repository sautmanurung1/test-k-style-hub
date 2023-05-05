package controller

import (
	"net/http"
	"skincare/interfaces"
	"skincare/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type reviewProductController struct {
	s interfaces.ReviewProductService
}

func NewReviewProductController(reviewProductService interfaces.ReviewProductService) interfaces.ReviewProductController {
	return &reviewProductController{
		s: reviewProductService,
	}
}

// InsertReview implements interfaces.ReviewProductController
func (h *reviewProductController) InsertReview(c echo.Context) error {
	reviews := models.ReviewProduct{
		ID: uuid.New().String(),
	}

	err := c.Bind(&reviews)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error to create review product",
			"message": err.Error(),
		})
	}

	fails := h.s.InsertReview(reviews)

	if fails != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error to insert data to database",
			"message": fails.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "success to insert review product to database",
	})
}
