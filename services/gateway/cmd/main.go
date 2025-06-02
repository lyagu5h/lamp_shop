package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/patrickmn/go-cache"
)

type OrderItem struct {
    ProductID int     `json:"product_id" validate:"required,min=1"`
    Quantity  int     `json:"quantity" validate:"required,gt=0"`
    Price     float64 `json:"price,omitempty"`
}

type Order struct {
    CustomerName  string      `json:"customer_name" validate:"required"`
    CustomerEmail string      `json:"customer_email" validate:"required,email"`
    CustomerPhone string      `json:"customer_phone" validate:"required"`
    Address       string      `json:"address" validate:"required"`
    Items         []OrderItem `json:"items" validate:"required,dive"`
    TotalAmount   float64     `json:"total_amount,omitempty"`
    Status        string      `json:"status,omitempty"`
}

type productResponse struct {
    ID          int             `json:"id"`
    Name        string          `json:"name"`
    Price       float64         `json:"price"`
    Power       int             `json:"power"`
    Description json.RawMessage `json:"description"`
    Temperature string          `json:"temperature"`
    Type        string          `json:"type"`
    LampCap     string          `json:"lamp_cap"`
    ImageURL    string          `json:"image_url"`
    CreatedAt   time.Time       `json:"created_at"`
    UpdatedAt   time.Time       `json:"updated_at"`
}

func main() {
    productsURL := os.Getenv("PRODUCTS_SERVICE_URL")
    ordersURL := os.Getenv("ORDERS_SERVICE_URL")
    if productsURL == "" || ordersURL == "" {
        log.Fatal("Need to set PRODUCTS_SERVICE_URL and ORDERS_SERVICE_URL")
    }


    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
	
    priceCache := cache.New(5*time.Minute, 10*time.Minute)

    app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    }))
    validate := validator.New()


    app.Post("/admin/products", func(c *fiber.Ctx) error {
        return proxy.Do(c, productsURL+"/products")
    })

    app.Get("/admin/orders", func(c *fiber.Ctx) error {
        return proxy.Do(c, ordersURL+"/orders")
    })


	app.Get("/products", func(c *fiber.Ctx) error {
		targetURL := fmt.Sprintf("%s%s", productsURL, c.OriginalURL())
		resp, err := http.Get(targetURL)
		if err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"products-service unreachable"})
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"read products response failed"})
		}

		var products []productResponse
		if err := json.Unmarshal(bodyBytes, &products); err != nil {
			c.Status(resp.StatusCode)
			c.Set("Content-Type", "application/json")
			return c.Send(bodyBytes)
		}

		for _, p := range products {
			key := fmt.Sprintf("price-%d", p.ID)
			priceCache.Set(key, p.Price, cache.DefaultExpiration)
		}

		c.Status(resp.StatusCode)
		c.Set("Content-Type", "application/json")
		return c.Send(bodyBytes)
		})

	app.Get("/products/:id", func(c *fiber.Ctx) error {
		targetURL := fmt.Sprintf("%s/products/%s", productsURL, c.Params("id"))
		resp, err := http.Get(targetURL)
		if err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"products-service unreachable"})
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"read product response failed"})
		}

		var pr productResponse
		if err := json.Unmarshal(bodyBytes, &pr); err == nil {
			key := fmt.Sprintf("price-%d", pr.ID)
			priceCache.Set(key, pr.Price, cache.DefaultExpiration)
		}
		c.Status(resp.StatusCode)
		c.Set("Content-Type", "application/json")
		return c.Send(bodyBytes)
	})

	app.All("/products/*", func(c *fiber.Ctx) error {
		target := fmt.Sprintf("%s%s", productsURL, c.OriginalURL())
		return proxy.Do(c, target)
	})


    app.All("/products/*", func(c *fiber.Ctx) error {
        target := fmt.Sprintf("%s%s", productsURL, c.OriginalURL())
        return proxy.Do(c, target)
    })

    app.Get("/orders", func(c *fiber.Ctx) error {
        target := fmt.Sprintf("%s/orders", ordersURL)
        return proxy.Do(c, target)
    })
    app.Get("/orders/:id", func(c *fiber.Ctx) error {
        target := fmt.Sprintf("%s/orders/%s", ordersURL, c.Params("id"))
        return proxy.Do(c, target)
    })
    app.Patch("/orders/:id/status", func(c *fiber.Ctx) error {
        target := fmt.Sprintf("%s/orders/%s/status", ordersURL, c.Params("id"))
        return proxy.Do(c, target)
    })

    app.Post("/orders", func(c *fiber.Ctx) error {
        var payload Order
        if err := c.BodyParser(&payload); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
        }
        if err := validate.Struct(&payload); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }
        if len(payload.Items) == 0 {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "items cannot be empty"})
        }

        total := 0.0
        for i := range payload.Items {
            pid := payload.Items[i].ProductID
            price, err := fetchProductPriceWithCache(c.Context(), priceCache, productsURL, pid)
            if err != nil {
                return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                    "error": fmt.Sprintf("cannot fetch product %d: %v", pid, err),
                })
            }
            payload.Items[i].Price = price
            total += price * float64(payload.Items[i].Quantity)
        }
        payload.TotalAmount = total
        payload.Status = "pending"

        bodyBytes, err := json.Marshal(payload)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "marshal failed"})
        }
        ordersCreateURL := fmt.Sprintf("%s/orders", ordersURL)
        req, err := http.NewRequestWithContext(c.Context(), http.MethodPost, ordersCreateURL, bytes.NewReader(bodyBytes))
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to build request"})
        }
        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{Timeout: 5 * time.Second}
        resp, err := client.Do(req)
        if err != nil {
            return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "orders-service unreachable"})
        }
        defer resp.Body.Close()

        respBody, err := io.ReadAll(resp.Body)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "read response failed"})
        }
        c.Status(resp.StatusCode)
        c.Set("Content-Type", "application/json")
        return c.Send(respBody)
    })

    log.Printf("Gateway is running on port %s (products: %s, orders: %s)\n", port, productsURL, ordersURL)
    log.Fatal(app.Listen(":" + port))
}

func fetchProductPriceWithCache(ctx context.Context, c *cache.Cache, productsURL string, productID int) (float64, error) {
    key := fmt.Sprintf("price-%d", productID)
    if x, found := c.Get(key); found {
        return x.(float64), nil
    }
    price, err := fetchProductPrice(ctx, productsURL, productID)
    if err != nil {
        return 0, err
    }
    c.Set(key, price, cache.DefaultExpiration)
    return price, nil
}

func fetchProductPrice(ctx context.Context, productsURL string, productID int) (float64, error) {
    url := fmt.Sprintf("%s/products/%d", productsURL, productID)
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return 0, err
    }
    client := &http.Client{Timeout: 4 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return 0, fmt.Errorf("status %d", resp.StatusCode)
    }
    var pr productResponse
    if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
        return 0, err
    }
    return pr.Price, nil
}
