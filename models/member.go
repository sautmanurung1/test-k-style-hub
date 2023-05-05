package models

import "gorm.io/plugin/soft_delete"

type Member struct {
	ID        string `json:"id" gorm:"primaryKey;size:255"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	SkinType  string `json:"skin_type"`
	SkinColor string `json:"skin_color"`

	CreatedAt int64                 `json:"-"`
	UpdatedAt int64                 `json:"-"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"index"`
}
