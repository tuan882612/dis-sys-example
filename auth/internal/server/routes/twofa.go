package routes

import (
	"github.com/labstack/echo/v4"

	"dissys/internal/auth/twofa"
	"dissys/internal/deps"
)

func SetTwoFARoutes(g *echo.Group, d *deps.Dependencies) error {
	svc, err := twofa.NewService(d)
	if err != nil {
		return err
	}

	handler := twofa.NewHandler(svc)

	twofaGroup := g.Group("/twofa")
	twofaGroup.POST("/login", handler.HandleLogin)
	twofaGroup.POST("/verify", handler.HandleVerify)
	twofaGroup.POST("/resend", handler.HandleResend)
	twofaGroup.POST("/register", handler.HandleRegister)
	twofaGroup.POST("/reset", handler.HandleReset)
	twofaGroup.POST("/reset/final", handler.HandleResetFinal)

	return nil
}
