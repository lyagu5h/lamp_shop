package httpapi

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/lyagu5h/auth-service/internal/usecase"
)

type AuthHandler struct {
    UC       *usecase.AuthUseCase
    Validate *validator.Validate
}

func NewAuthHandler(uc *usecase.AuthUseCase, v *validator.Validate) *AuthHandler {
    return &AuthHandler{UC: uc, Validate: v}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
    var req struct {
        Username string `json:"username" validate:"required,min=3"`
        Password string `json:"password" validate:"required,min=8"`
        Role     string `json:"role" validate:"required,oneof=admin catalog_manager order_manager"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
    }
    if err := h.Validate.Struct(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    if err := h.UC.Register(req.Username, req.Password, req.Role); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create user"})
    }
    return c.SendStatus(http.StatusCreated)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var req struct {
        Username string `json:"username" validate:"required"`
        Password string `json:"password" validate:"required"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
    }
    if err := h.Validate.Struct(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    token, err := h.UC.Login(req.Username, req.Password)
    if err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
    }
    return c.JSON(fiber.Map{
        "token":      token,
        "expires_at": time.Now().Add(time.Hour * 8).Unix(),
    })
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
    claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
    return c.JSON(fiber.Map{
        "sub":  claims["sub"],
        "role": claims["role"],
        "exp":  claims["exp"],
    })
}
