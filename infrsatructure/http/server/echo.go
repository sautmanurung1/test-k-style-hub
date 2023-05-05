package server

import (
	"net/http"
	"skincare/infrsatructure/database"
	"skincare/infrsatructure/http/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server() *echo.Echo {
	app := echo.New()
	conf := database.Config{}

	routes.RootRoutes(app, conf)

	app.Use(middleware.CORS())
	//app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": true,
		})
	})

	return app
}
