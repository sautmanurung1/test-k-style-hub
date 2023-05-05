package routes

import (
	c "skincare/controller"
	"skincare/infrsatructure/database"
	r "skincare/repository"
	s "skincare/service"

	"github.com/labstack/echo/v4"
)

func RoutesMember(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	memberRepository := r.NewMemberRepository(db)
	memberService := s.NewMemberService(memberRepository, conf)
	memberController := c.NewMemberController(memberService)

	member := echo.Group("/v1/member")
	member.POST("/create", memberController.InsertMember)
	member.GET("/all", memberController.GetAllMember)
	member.PUT("/:id", memberController.UpdateMember)
	member.DELETE("/:id", memberController.DeleteMember)
}
