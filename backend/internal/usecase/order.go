package usecase

import "github.com/lyagu5h/lamp_backend/internal/domain"

type OrderUseCase struct {
	Repo domain.OrderRepository
}

func NewOrderUseCase(r domain.OrderRepository) *OrderUseCase {
	return &OrderUseCase{Repo: r}
}

func (uc *OrderUseCase) List() ([]domain.Order, error) {
	return uc.Repo.List()
}
func (uc *OrderUseCase) Get(id int) (domain.Order, error) {
	return uc.Repo.GetByID(id)
}
func (uc *OrderUseCase) Create(o *domain.Order) error {
	return uc.Repo.Create(o)
}
func (uc *OrderUseCase) UpdateStatus(id int, status string) error {
	return uc.Repo.UpdateStatus(id, status)
}