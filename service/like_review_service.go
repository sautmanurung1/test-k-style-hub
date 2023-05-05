package service

import (
	"skincare/infrsatructure/database"
	"skincare/interfaces"
	"skincare/models"
)

type likeReviewService struct {
	c    database.Config
	repo interfaces.LikeReviewRepository
}

func NewLikeReviewService(repo interfaces.LikeReviewRepository, c database.Config) interfaces.LikeReviewService {
	return &likeReviewService{
		repo: repo,
		c:    c,
	}
}

// DeleteLikeReview implements interfaces.LikeReviewService
func (s *likeReviewService) DeleteLikeReview(id string, like models.LikeReview) {
	s.repo.DeleteLikeReview(id, like)
}

// InsertLikeReview implements interfaces.LikeReviewService
func (s *likeReviewService) InsertLikeReview(like models.LikeReview) error {
	return s.repo.InsertLikeReview(like)
}
