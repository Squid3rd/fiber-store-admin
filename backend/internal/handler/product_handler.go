package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/squid3rd/fiber-store-admin/internal/model"
	"github.com/squid3rd/fiber-store-admin/internal/service"
	"github.com/squid3rd/fiber-store-admin/pkg/response"
)

type ProductHandler struct {
	service *service.ProductService
}

// type HealthHandler struct {
// 	service *service.HealthService
// }

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// FindAll godoc
// @Summary List products
// @Description Paginated list of products
// @Tags products
// @Produce json
// @Param page query int false "page number" default(1)
// @Param limit query int false "page size" default(20)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (h *ProductHandler) FindAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "20"))

	items, total, err := h.service.FindAll(c.Context(), int64(page), int64(limit))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": items,
		"meta": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// Create godoc
// @Summary Create product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param payload body model.CreateProductInput true "product payload"
// @Success 200 {object} model.Product
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [post]
func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var in model.CreateProductInput

	if err := c.BodyParser(&in); err != nil {
		return response.Fail(c, "INVALID_REQUEST_BODY", err.Error())
	}

	if in.Name == "" {
		return response.Fail(c, "INVALID_REQUEST_BODY", "Name is required")
	}

	p, err := h.service.Create(c.Context(), in)
	if err != nil {
		return response.Fail(c, "PRODUCT_CREATE_ERROR", err.Error())
	}

	return response.OK(c, p)
}

// func (h *ProductHandler) HealthHandler(c *fiber.Ctx) error {
// 	return response.OK(c, "OK")
// }
