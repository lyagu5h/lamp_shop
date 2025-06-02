package main

import (
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"

	"github.com/lyagu5h/auth-service/internal/infrastructure/db"
	"github.com/lyagu5h/auth-service/internal/interface/httpapi"
	"github.com/lyagu5h/auth-service/internal/repository"
	"github.com/lyagu5h/auth-service/internal/usecase"
)

func main() {
    _ = godotenv.Load()

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET не задан в .env или окружении")
    }

    pg := db.NewPostgresDB()
    defer pg.Close()
    if err := db.RunMigrations(pg, "./migrations"); err != nil {
        log.Fatalf("Ошибка миграций: %v", err)
    }

    authUC := usecase.NewAuthUseCase(repository.NewUserRepo(pg), []byte(jwtSecret), 8*time.Hour)
    handler := httpapi.NewAuthHandler(authUC, validator.New())

    app := fiber.New()
    app.Use(cors.New())

    app.Post("/auth/register", handler.Register)
    app.Post("/auth/login", handler.Login)

    protected := jwtware.New(jwtware.Config{
        SigningKey:    []byte(jwtSecret),
        ContextKey:    "user",
        TokenLookup:   "header:Authorization",
        AuthScheme:    "Bearer",
        SigningMethod: "HS256",
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": err.Error(),
            })
        },
    })
    app.Get("/auth/me", protected, handler.Me)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8082"
    }
    log.Println("Auth-service стартует на порту:", port)
    log.Fatal(app.Listen(":" + port))
}
