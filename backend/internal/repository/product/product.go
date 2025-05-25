package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lyagu5h/lamp_backend/internal/domain"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) domain.ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) GetAll() ([]domain.Product, error) {
	var pArr []domain.Product 

	err := r.db.Select(&pArr, "SELECT * FROM products ORDER BY id")

	return pArr, err
}

func (r *productRepo) GetByID(id int) (domain.Product, error) {
	var p domain.Product
	err := r.db.Get(&p, "SELECT * FROM products WHERE id=$1", id)

	return p, err
}

func (r *productRepo) Create(p *domain.Product) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return r.db.QueryRowx(
		`INSERT INTO products (name, price, power, temperature, type, lamp_cap, image_url, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		p.Name, p.Price, p.Power, p.Temperature, p.Type, p.LampCap, p.ImageURL, p.CreatedAt, p.UpdatedAt,
		).Scan(&p.ID)
}

func (r *productRepo) Update(p *domain.Product) error {
    p.UpdatedAt = time.Now()
    _, err := r.db.Exec(
        `UPDATE products SET name=$1,price=$2,power=$3,temperature=$4,type=$5,lamp_cap=$6,image_url=$7,updated_at=$8 WHERE id=$9`,
        p.Name,p.Price,p.Power,p.Temperature,p.Type,p.LampCap,p.ImageURL,p.UpdatedAt,p.ID,
    )
    return err
}

func (r *productRepo) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
    return err
}