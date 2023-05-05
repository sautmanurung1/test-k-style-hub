package models

import "gorm.io/plugin/soft_delete"

type Product struct {
	ID          string                `json:"id" gorm:"primaryKey;size:255"`
	ProductName string                `json:"product_name"`
	Price       int                   `json:"price"`
	CreatedAt   int64                 `json:"-"`
	UpdatedAt   int64                 `json:"-"`
	DeletedAt   soft_delete.DeletedAt `json:"-" gorm:"index"`
}
