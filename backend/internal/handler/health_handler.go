package handler

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

// Check health of the API
// @Summary Check health of the API
// @Description Check health of the API
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /health [get]
func (h *HealthHandler) Check(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
