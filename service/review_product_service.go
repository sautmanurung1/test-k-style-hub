package service

import (
	"skincare/infrsatructure/database"
	"skincare/interfaces"
	"skincare/models"
)

type reviewProductService struct {
	c    database.Config
	repo interfaces.ReviewProductRepository
}

func NewReviewProductService(repo interfaces.ReviewProductRepository, c database.Config) interfaces.ReviewProductService {
	return &reviewProductService{
		repo: repo,
		c:    c,
	}
}

// InsertReview implements interfaces.ReviewProductService
func (s *reviewProductService) InsertReview(review models.ReviewProduct) error {
	return s.repo.InsertReview(review)
}
