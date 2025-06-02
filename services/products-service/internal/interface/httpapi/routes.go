package httpapi

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lyagu5h/products-service/internal/usecase"
	"github.com/minio/minio-go/v7"
)

func RegisterRoutes(app *fiber.App, uc *usecase.ProductUseCase,
                    v *validator.Validate, minioCli *minio.Client, bucket string,
                    auth fiber.Handler, admin fiber.Handler) {

    handler := NewProductHandler(uc, v, minioCli, bucket)
    grpPublic := app.Group("/products")

    grpPublic.Get("/", handler.List)
    grpPublic.Get("/:id", handler.Get)

	grpProtected := grpPublic.Group("", auth)

    grpProtected.Post("/", handler.Create)
    grpProtected.Put("/:id", handler.Update)
    grpProtected.Delete("/:id", handler.Delete)
    grpProtected.Post("/:id/image", handler.uploadProductImage)
}