package repository

import (
	"skincare/interfaces"
	"skincare/models"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productRepository{
		DB: db,
	}
}

// InsertProduct implements interfaces.ProductRepository
func (r *productRepository) InsertProduct(product models.Product) error {
	response := r.DB.Create(&product)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

// GetProduct implements interfaces.ProductRepository
func (r *productRepository) GetProduct(id string) (review models.ReviewProduct, result int64, err error) {
	var likeReview models.LikeReview
	r.DB.Find(&likeReview).Count(&result)
	r.DB.Where("product_id = ?", id).Preload("Member").Preload("Product").Find(&review)
	return review, result, nil
}
