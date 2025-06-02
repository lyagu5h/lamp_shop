package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/lyagu5h/auth-service/internal/repository"
)

type AuthUseCase struct {
    userRepo   *repository.UserRepo
    jwtSecret  []byte
    jwtExpires time.Duration
}

func NewAuthUseCase(repo *repository.UserRepo, secret []byte, expires time.Duration) *AuthUseCase {
    return &AuthUseCase{userRepo: repo, jwtSecret: secret, jwtExpires: expires}
}

func (uc *AuthUseCase) Register(username, password, role string) error {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user := &repository.User{
        Username:     username,
        PasswordHash: string(hash),
        Role:         role,
    }
    return uc.userRepo.Create(user)
}

func (uc *AuthUseCase) Login(username, password string) (string, error) {
    user, err := uc.userRepo.GetByUsername(username)
    if err != nil {
        return "", errors.New("invalid credentials")
    }
    if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
        return "", errors.New("invalid credentials")
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub":  user.Username,
        "role": user.Role,
        "exp":  time.Now().Add(uc.jwtExpires).Unix(),
    })
    signed, err := token.SignedString(uc.jwtSecret)
    if err != nil {
        return "", err
    }
    return signed, nil
}

