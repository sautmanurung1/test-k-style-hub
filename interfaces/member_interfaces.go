package interfaces

import (
	"skincare/models"

	"github.com/labstack/echo/v4"
)

type MemberRepository interface {
	InsertMember(member models.Member) error
	UpdateMember(id string, member models.Member)
	DeleteMember(id string, member models.Member)
	GetAllMember() (member []models.Member, err error)
}

type MemberService interface {
	InsertMember(member models.Member) error
	UpdateMember(id string, member models.Member)
	DeleteMember(id string, member models.Member)
	GetAllMember() (member []models.Member, err error)
}

type MemberController interface {
	InsertMember(c echo.Context) error
	UpdateMember(c echo.Context) error
	DeleteMember(c echo.Context) error
	GetAllMember(c echo.Context) error
}
