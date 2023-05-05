package controller

import (
	"net/http"
	"skincare/interfaces"
	"skincare/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type memberController struct {
	s interfaces.MemberService
}

func NewMemberController(membersService interfaces.MemberService) interfaces.MemberController {
	return &memberController{
		s: membersService,
	}
}

// DeleteMember implements interfaces.MemberController
func (h *memberController) DeleteMember(c echo.Context) error {
	var member models.Member
	id := c.Param("id")
	if err := c.Bind(&member); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error to delete id member",
			"message": err.Error(),
		})
	}

	h.s.DeleteMember(id, member)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success to delete member from id",
		"code":   http.StatusOK,
	})
}

// GetAllMember implements interfaces.MemberController
func (h *memberController) GetAllMember(c echo.Context) error {
	member, err := h.s.GetAllMember()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error to get all member",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to get all member",
		"data":    member,
	})
}

// InsertMember implements interfaces.MemberController
func (h *memberController) InsertMember(c echo.Context) error {
	members := models.Member{
		ID: uuid.New().String(),
	}

	err := c.Bind(&members)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error to create member",
			"message": err.Error(),
		})
	}

	fails := h.s.InsertMember(members)

	if fails != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error to insert data to database",
			"message": fails.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "success to insert member to database",
	})
}

// UpdateMember implements interfaces.MemberController
func (h *memberController) UpdateMember(c echo.Context) error {
	var member models.Member
	id := c.Param("id")

	if err := c.Bind(&member); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error to update member data",
			"status":  err.Error(),
		})
	}

	h.s.UpdateMember(id, member)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success to update member from id",
		"code":   http.StatusOK,
	})
}
