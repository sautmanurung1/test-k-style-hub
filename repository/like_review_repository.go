package repository

import (
	"skincare/interfaces"
	"skincare/models"

	"gorm.io/gorm"
)

type likeReviewRepository struct {
	DB *gorm.DB
}

func NewLikeReviewRepository(db *gorm.DB) interfaces.LikeReviewRepository {
	return &likeReviewRepository{
		DB: db,
	}
}

// DeleteLikeReview implements interfaces.LikeReviewRepository
func (r *likeReviewRepository) DeleteLikeReview(id string, like models.LikeReview) {
	r.DB.Where("id = ?", id).Delete(&like)
}

// InsertLikeReview implements interfaces.LikeReviewRepository
func (r *likeReviewRepository) InsertLikeReview(like models.LikeReview) error {
	response := r.DB.Create(&like)
	if response.Error != nil {
		return response.Error
	}
	return nil
}
