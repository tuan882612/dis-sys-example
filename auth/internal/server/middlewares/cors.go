package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetCORS(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost:3000"},
		AllowHeaders: []string{
			echo.HeaderAuthorization,
			"X-Uid",
			"X-API-Key",
			"X-Request-Type",
			"X-Validation-Type",
		},
	}))
}
