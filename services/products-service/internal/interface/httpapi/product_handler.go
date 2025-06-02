package httpapi

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lyagu5h/products-service/internal/domain"
	"github.com/lyagu5h/products-service/internal/usecase"
	"github.com/minio/minio-go/v7"
)

type ProductHandler struct {
	UC *usecase.ProductUseCase
	Validate *validator.Validate
	MinioCli *minio.Client
	Bucket string
}

func NewProductHandler(uc *usecase.ProductUseCase, validate *validator.Validate, minioCli *minio.Client, bucket string) *ProductHandler {
	return &ProductHandler{UC: uc, Validate: validate, MinioCli: minioCli, Bucket: bucket}
}

func (h *ProductHandler) List(c *fiber.Ctx) error {
	items, err := h.UC.List()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(items)
}

func (h *ProductHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	item, err := h.UC.Get(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "product not found"})
	}

	return c.JSON(item)	
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var p domain.Product
	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	if err := h.Validate.Struct(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = p.CreatedAt

	if err := h.UC.Create(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(p)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var p domain.Product
	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	p.ID = id
	if err := h.UC.Update(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "success"})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.UC.Delete(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted successfully"})
}

func (h *ProductHandler) uploadProductImage(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product id"})
    }
    fileHeader, ferr := c.FormFile("image")
    if ferr != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no image file"})
    }
    file, err := fileHeader.Open()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot open uploaded file"})
    }
    defer file.Close()

    filename := fmt.Sprintf("%d_%d_%s", id, time.Now().Unix(), fileHeader.Filename)

    ctx := context.Background()

	contentType := fileHeader.Header.Get("Content-Type")
    uploadInfo, err := h.MinioCli.PutObject(ctx, h.Bucket, filename, file, fileHeader.Size, minio.PutObjectOptions{
        ContentType: contentType,
    })
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to upload to storage"})
    }

    imageURL := fmt.Sprintf("%s/%s/%s", "http://localhost:9000", h.Bucket, uploadInfo.Key)

    if err := h.UC.UpdateImageURL(c.Context(), id, imageURL); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update image url in db"})
    }

    return c.JSON(fiber.Map{"image_url": imageURL})
}
