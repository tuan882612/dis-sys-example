package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetBaseRoutes(g *echo.Group) {
	g.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "OK",
			"service": "auth",
		},
		)
	})
}
