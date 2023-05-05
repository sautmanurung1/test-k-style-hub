package service

import (
	"skincare/infrsatructure/database"
	"skincare/interfaces"
	"skincare/models"
)

type svcMember struct {
	c    database.Config
	repo interfaces.MemberRepository
}

func NewMemberService(repo interfaces.MemberRepository, c database.Config) interfaces.MemberService {
	return &svcMember{
		repo: repo,
		c:    c,
	}
}

// DeleteMember implements interfaces.MemberService
func (s *svcMember) DeleteMember(id string, member models.Member) {
	s.repo.DeleteMember(id, member)
}

// GetAllMember implements interfaces.MemberService
func (s *svcMember) GetAllMember() (member []models.Member, err error) {
	return s.repo.GetAllMember()
}

// InsertMember implements interfaces.MemberService
func (s *svcMember) InsertMember(member models.Member) error {
	return s.repo.InsertMember(member)
}

// UpdateMember implements interfaces.MemberService
func (s *svcMember) UpdateMember(id string, member models.Member) {
	s.repo.UpdateMember(id, member)
}
