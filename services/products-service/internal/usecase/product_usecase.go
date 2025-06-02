package usecase

import (
	"context"

	"github.com/lyagu5h/products-service/internal/domain"
)

type ProductUseCase struct {
	productRepo domain.ProductRepository
}


func NewProductUseCase(productRepo domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepo: productRepo}
}

func (u *ProductUseCase) List() ([]domain.Product, error) {
	return u.productRepo.GetAll()
}

func (u *ProductUseCase) Get(id int) (domain.Product, error) {
	return u.productRepo.GetByID(id)
}

func (u *ProductUseCase) Create(p *domain.Product) error {
	return u.productRepo.Create(p)
}

func (u *ProductUseCase) Update(p *domain.Product) error {
	return u.productRepo.Update(p)
}

func (u *ProductUseCase) Delete(id int) error {
	return u.productRepo.Delete(id)
}

func (u *ProductUseCase) UpdateImageURL(ctx context.Context, id int, url string) error {
    return u.productRepo.UpdateImageURL(ctx, id, url)
}


