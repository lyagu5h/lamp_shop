package httpapi

import (
	"strconv"

	"github.com/lyagu5h/lamp_backend/internal/domain"
	"github.com/lyagu5h/lamp_backend/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	UC       *usecase.OrderUseCase
	Validate *validator.Validate
}

func NewOrderHandler(uc *usecase.OrderUseCase, v *validator.Validate) *OrderHandler {
	return &OrderHandler{UC: uc, Validate: v}
}

func (h *OrderHandler) List(c *fiber.Ctx) error {
	orders, err := h.UC.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(orders)
}

func (h *OrderHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	o, err := h.UC.Get(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "order not found"})
	}
	return c.JSON(o)
}

func (h *OrderHandler) Create(c *fiber.Ctx) error {
	var o domain.Order
	if err := c.BodyParser(&o); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	if err := h.Validate.Struct(&o); err != nil {
		return err
	}
	if err := h.UC.Create(&o); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(o)
}

func (h *OrderHandler) UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var payload struct {
		Status string `json:"status" validate:"required"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	if err := h.Validate.Var(payload.Status, "required"); err != nil {
		return err
	}
	if err := h.UC.UpdateStatus(id, payload.Status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Order status updated"})
}