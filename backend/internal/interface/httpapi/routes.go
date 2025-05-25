package httpapi

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lyagu5h/lamp_backend/internal/usecase"
)

func RegisterRoutes(app *fiber.App, productUC *usecase.ProductUseCase, orderUC *usecase.OrderUseCase, validate *validator.Validate) {
	handler := NewProductHandler(productUC, validate)
	grp := app.Group("/products")
	grp.Get("/", handler.List)
	grp.Get("/:id", handler.Get)
	grp.Post("/", handler.Create)
	grp.Put("/:id", handler.Update)
	grp.Delete("/:id", handler.Delete)

	grp = app.Group("/orders")
	oh := NewOrderHandler(orderUC, validate)
	grp.Get("/", oh.List)
	grp.Get("/:id", oh.Get)
	grp.Post("/", oh.Create)
	grp.Patch("/:id/status", oh.UpdateStatus)
}