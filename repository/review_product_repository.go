package repository

import (
	"skincare/interfaces"
	"skincare/models"

	"gorm.io/gorm"
)

type reviewProductRepository struct {
	DB *gorm.DB
}

func NewReviewProductRepository(db *gorm.DB) interfaces.ReviewProductRepository {
	return &reviewProductRepository{
		DB: db,
	}
}

// InsertReview implements interfaces.ReviewProductRepository
func (r *reviewProductRepository) InsertReview(review models.ReviewProduct) error {
	response := r.DB.Create(&review)
	if response != nil {
		return response.Error
	}
	return nil
}
