package service

import (
	"skincare/infrsatructure/database"
	"skincare/interfaces"
	"skincare/models"
)

type productService struct {
	c    database.Config
	repo interfaces.ProductRepository
}

func NewProductService(repo interfaces.ProductRepository, c database.Config) interfaces.ProductService {
	return &productService{
		repo: repo,
		c:    c,
	}
}

// InsertProduct implements interfaces.ProductService
func (s *productService) InsertProduct(product models.Product) error {
	return s.repo.InsertProduct(product)
}

// GetProduct implements interfaces.ProductService
func (s *productService) GetProduct(id string) (review models.ReviewProduct, result int64, err error) {
	return s.repo.GetProduct(id)
}
