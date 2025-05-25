package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	infra "github.com/lyagu5h/lamp_backend/internal/infrastructure/db"
	"github.com/lyagu5h/lamp_backend/internal/interface/httpapi"
	orderRepo "github.com/lyagu5h/lamp_backend/internal/repository/order"
	productRepo "github.com/lyagu5h/lamp_backend/internal/repository/product"
	"github.com/lyagu5h/lamp_backend/internal/usecase"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	db, err := infra.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	if err := infra.RunMigration(db, "./migrations"); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	validate := validator.New()

	pRepo := productRepo.NewProductRepo(db)
	pUC := usecase.NewProductUseCase(pRepo)

	oRepo := orderRepo.NewOrderRepo(db)
	orderUC := usecase.NewOrderUseCase(oRepo)

	app := fiber.New(fiber.Config{
		ErrorHandler: func (c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msg := "internal server error"
			if ve, ok := err.(validator.ValidationErrors); ok {
				code = fiber.StatusBadRequest
				msg = ve.Error()
			}
			return c.Status(code).JSON(fiber.Map{"error": msg})
		},
	})

	httpapi.RegisterRoutes(app, pUC, orderUC, validate)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
	log.Printf("Server listening on: %s", port)
}