package models

import "gorm.io/plugin/soft_delete"

type ReviewProduct struct {
	ID string `json:"id" gorm:"primaryKey;size:255"`

	// Foreign Key
	ProductId string `json:"product_id" gorm:"size:255"`
	MemberId  string `json:"member_id" gorm:"size:255"`

	DescReview string `json:"desc_review"`

	// Relationship
	Member  Member  `json:"member" gorm:"foreignkey:MemberId;references:ID"`
	Product Product `json:"product" gorm:"foreignkey:ProductId;references:ID"`

	CreatedAt int64                 `json:"-"`
	UpdatedAt int64                 `json:"-"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"index"`
}
