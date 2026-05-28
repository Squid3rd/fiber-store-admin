package response

import "github.com/gofiber/fiber/v2"

type Success struct {
	Data any   `json:"data"`
	Meta *Meta `json:"meta"`
}

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func OK(c *fiber.Ctx, data any) error {
	return c.JSON(fiber.Map{"data": Success{Data: data}})
}

func Fail(c *fiber.Ctx, code string, message string) error {
	return c.JSON(fiber.Map{"error": ErrorBody{Code: code, Message: message}})
}
