package repository

import (
	"skincare/interfaces"
	"skincare/models"

	"gorm.io/gorm"
)

type memberRepository struct {
	DB *gorm.DB
}

func NewMemberRepository(db *gorm.DB) interfaces.MemberRepository {
	return &memberRepository{
		DB: db,
	}
}

// DeleteMember implements interfaces.MemberRepository
func (r *memberRepository) DeleteMember(id string, member models.Member) {
	r.DB.Where("id = ?", id).Delete(&member)
}

// GetAllMember implements interfaces.MemberRepository
func (r *memberRepository) GetAllMember() (member []models.Member, err error) {
	r.DB.Find(&member)
	return member, nil
}

// InsertMember implements interfaces.MemberRepository
func (r *memberRepository) InsertMember(member models.Member) error {
	response := r.DB.Create(&member)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

// UpdateMember implements interfaces.MemberRepository
func (r *memberRepository) UpdateMember(id string, member models.Member) {
	r.DB.Model(&models.Member{}).Where("id", id).Updates(member)
}
