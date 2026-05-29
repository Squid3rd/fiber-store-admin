package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/squid3rd/fiber-store-admin/internal/model"
	"github.com/squid3rd/fiber-store-admin/internal/service"
	"github.com/squid3rd/fiber-store-admin/pkg/response"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) CreateUser(c *fiber.Ctx) error {
	var in model.User

	if err := c.BodyParser(&in); err != nil {
		return response.Fail(c, "INVALID_REQUEST_BODY", err.Error())
	}

	if in.Name == "" {
		return response.Fail(c, "INVALID_REQUEST_BODY", "Name is required")
	}

	if in.Email == "" {
		return response.Fail(c, "INVALID_REQUEST_BODY", "Email is required")
	}

	if in.Password == "" {
		return response.Fail(c, "INVALID_REQUEST_BODY", "Password is required")
	}

	u, err := h.service.CreateUser(c.Context(), in)
	if err != nil {
		return response.Fail(c, "USER_CREATE_ERROR", err.Error())
	}

	return response.OK(c, u)
}
