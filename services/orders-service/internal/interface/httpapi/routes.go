package httpapi

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lyagu5h/orders-service/internal/usecase"
)

func RegisterRoutes(app *fiber.App, orderUC *usecase.OrderUseCase, validate *validator.Validate, auth fiber.Handler, admin fiber.Handler) {
    handler := NewOrderHandler(orderUC, validate)
    grpPublic := app.Group("/orders")
    grpPublic.Get("/", handler.List)
    grpPublic.Get("/:id", handler.Get)
    grpPublic.Post("/", handler.Create)

    grpProtected := grpPublic.Group("", auth)
    grpProtected.Patch("/:id/status", handler.UpdateStatus)
}
