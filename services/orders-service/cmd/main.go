package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"

	db "github.com/lyagu5h/orders-service/internal/infrastructure/db"
	"github.com/lyagu5h/orders-service/internal/interface/httpapi"
	orderRepo "github.com/lyagu5h/orders-service/internal/repository"
	"github.com/lyagu5h/orders-service/internal/usecase"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("warning: .env file not found or could not be loaded")
    }

	dbConn := db.NewPostgresDB()
	defer dbConn.Close()

	if err := db.RunMigration(dbConn, "./migrations"); err != nil {
		log.Fatalf("goose up (orders) failed: %v", err)
	}
    
    orderRepo := orderRepo.NewOrderRepo(dbConn)
    orderUC := usecase.NewOrderUseCase(orderRepo)

    app := fiber.New()
    validate := validator.New()

	app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    }))


    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET не задан в окружении orders-service")
    }

    authMW := jwtware.New(jwtware.Config{
        SigningKey:  []byte(jwtSecret),
        ContextKey:  "user",
        TokenLookup: "header:Authorization",
        AuthScheme:  "Bearer",
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
        },
    })

    adminOnly := func(c *fiber.Ctx) error {
        token := c.Locals("user").(*jwt.Token)
        claims := token.Claims.(jwt.MapClaims)
        if claims["role"] != "admin" {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Admin rights required"})
        }
        return c.Next()
    }

    httpapi.RegisterRoutes(app, orderUC, validate, authMW, adminOnly)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8090"
    }
    log.Printf("orders-service is running on port %s", port)
    log.Fatal(app.Listen(":" + port))
}