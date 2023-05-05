package models

import "gorm.io/plugin/soft_delete"

type LikeReview struct {
	ID string `json:"id" gorm:"primaryKey"`

	// Foreign Key
	ReviewId  string `json:"review_id" gorm:"size:255"`
	MembersId string `json:"member_id" gorm:"size:255"`

	// Relantionship
	ReviewProduct ReviewProduct `json:"review" gorm:"foreignKey:ReviewId;references:ID"`
	Member        Member        `json:"member" gorm:"foreignKey:MembersId;references:ID"`

	CreatedAt int64                 `json:"-"`
	UpdatedAt int64                 `json:"-"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"index"`
}
