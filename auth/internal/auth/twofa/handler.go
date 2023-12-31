package twofa

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"dissys/internal/auth"
	"dissys/pkg/httputils"
)

type Handler struct {
	service *tfaService
}

func NewHandler(twofaSvc *tfaService) *Handler {
	return &Handler{
		service: twofaSvc,
	}
}

func (h *Handler) HandleLogin(c echo.Context) error {
	req := auth.Request{}
	if err := httputils.Unmarshal(c.Request().Body, &req); err != nil {
		return err
	}

	userID, err := h.service.Login(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	c.Response().Header().Set("X-Request-Type", "login")
	c.Response().Header().Set("X-Uid", userID.String())
	return c.JSON(200, nil)
}

func (h *Handler) HandleVerify(c echo.Context) error {
	reqType := c.Request().Header.Get("X-Request-Type")
	if reqType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("missing header: %s", "X-Request-Type"),
		})
	}

	return nil
}

func (h *Handler) HandleResend(c echo.Context) error {
	return nil
}

func (h *Handler) HandleRegister(c echo.Context) error {
	req := &auth.Request{}
	if err := httputils.Unmarshal(c.Request().Body, req); err != nil {
		return err
	}

	userID, err := h.service.Register(c.Request().Context(), req)
	if err != nil {
		return err
	}

	c.Response().Header().Set("X-Request-Type", "register")
	c.Response().Header().Set("X-Uid", userID.String())
	return c.JSON(200, nil)
}

func (h *Handler) HandleReset(c echo.Context) error {
	return nil
}

func (h *Handler) HandleResetFinal(c echo.Context) error {
	return nil
}
