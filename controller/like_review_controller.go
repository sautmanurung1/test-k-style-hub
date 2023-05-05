package controller

import (
	"net/http"
	"skincare/interfaces"
	"skincare/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type likeReviewController struct {
	s interfaces.LikeReviewService
}

func NewLikeReviewController(likeReviewService interfaces.LikeReviewService) interfaces.LikeReviewController {
	return &likeReviewController{
		s: likeReviewService,
	}
}

// DeleteLikeReview implements interfaces.LikeReviewController
func (h *likeReviewController) DeleteLikeReview(c echo.Context) error {
	var like models.LikeReview
	id := c.Param("id")
	if err := c.Bind(&like); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error to delete id like review",
			"message": err.Error(),
		})
	}

	h.s.DeleteLikeReview(id, like)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success to dislike the review",
		"code":   http.StatusOK,
	})
}

// InsertLikeReview implements interfaces.LikeReviewController
func (h *likeReviewController) InsertLikeReview(c echo.Context) error {
	likes := models.LikeReview{
		ID: uuid.New().String(),
	}

	if err := c.Bind(&likes); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error to like the review",
			"message": err.Error(),
		})
	}

	fails := h.s.InsertLikeReview(likes)

	if fails != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error to like the review",
			"message": fails.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "success to give like the review",
	})
}
