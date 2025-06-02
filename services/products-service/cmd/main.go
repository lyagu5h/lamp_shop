package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	db "github.com/lyagu5h/products-service/internal/infrastructure/db"
	min "github.com/lyagu5h/products-service/internal/infrastructure/minio"
	"github.com/lyagu5h/products-service/internal/interface/httpapi"
	"github.com/lyagu5h/products-service/internal/repository"
	"github.com/lyagu5h/products-service/internal/usecase"
	"github.com/minio/minio-go/v7"
	"github.com/pressly/goose/v3"
)



func main() {
    _ = godotenv.Load()

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET не задан в .env или окружении products-service")
    }

    pg := db.NewPostgresDB()
    defer pg.Close()
    productRepo := repository.NewProductRepo(pg)
    productUC := usecase.NewProductUseCase(productRepo)

   
    minioClient := min.NewMinioClient()
    bucket := os.Getenv("MINIO_BUCKET_PRODUCTS")
    if bucket == "" {
        log.Fatal("MINIO_BUCKET_PRODUCTS must be set")
    }

    migrationsDir := "./migrations"
    if err := goose.SetDialect("postgres"); err != nil {
        log.Fatalf("goose set dialect: %v", err)
    }
    if err := goose.Up(pg.DB, migrationsDir); err != nil {
        log.Fatalf("goose up failed: %v", err)
    }
    if err := seedProducts(pg, minioClient, bucket); err != nil {
        log.Fatalf("seeding products failed: %v", err)
    }
    app := fiber.New()
    app.Use(cors.New())

    authMW := jwtware.New(jwtware.Config{
        SigningKey:    []byte(jwtSecret),
        ContextKey:    "user",
        TokenLookup:   "header:Authorization",
        AuthScheme:    "Bearer",
        SigningMethod: "HS256",
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

    httpapi.RegisterRoutes(app, productUC, validator.New(), minioClient, bucket, authMW, adminOnly)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("products-service запущен на порту: %s\n", port)
    log.Fatal(app.Listen(":" + port))
}

func seedProducts(dbConn *sqlx.DB, minioCli *minio.Client, bucket string) error {

    imagesDir := "seeds/seed_images"
    files, err := os.ReadDir(imagesDir)
    if err != nil {
        return err
    }

    
    ctx := context.Background()

    for _, entry := range files {
        if entry.IsDir() {
            continue
        }
        filename := entry.Name() 
        ext := filepath.Ext(filename)            
        base := filename[0 : len(filename)-len(ext)] 
        var pid int
        _, err := fmt.Sscanf(base, "%d", &pid)
        if err != nil {
            log.Printf("cannot parse product id from filename %s: %v; skipping", filename, err)
            continue
        }
        filePath := filepath.Join(imagesDir, filename)
        f, err := os.Open(filePath)
        if err != nil {
            return err
        }
        defer f.Close()

        objectName := "products/" + filename
        info, err := minioCli.PutObject(ctx, bucket, objectName, f, -1, minio.PutObjectOptions{
            ContentType: "image/jpeg",
        })
        if err != nil {
            return err
        }
        log.Printf("Uploaded seed image %s, size %d\n", objectName, info.Size)

        imageURL := fmt.Sprintf("http://%s/%s/%s", "localhost:9000", bucket, objectName)


        _, err = dbConn.Exec(`
            UPDATE products
            SET image_url = $1,
                updated_at = $2
            WHERE id = $3
        `, imageURL, time.Now().UTC(), pid)
        if err != nil {
            return err
        }
    }

    return nil
}

