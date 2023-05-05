package routes

import (
	"skincare/infrsatructure/database"

	"github.com/labstack/echo/v4"
)

func RootRoutes(echo *echo.Echo, conf database.Config) {
	RoutesMember(echo, conf)
	RoutesProduct(echo, conf)
	ReviewProductRoutes(echo, conf)
	LikeReviewRoutes(echo, conf)
}
