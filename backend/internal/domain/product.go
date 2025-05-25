package domain

import (
	"database/sql"
	"time"
)

type Product struct {
	ID int `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
    Price       float64   `db:"price" json:"price"`
    Power       int       `db:"power" json:"power"`
    Description sql.NullString    `db:"description" json:"description"`
    Temperature string    `db:"temperature" json:"temperature"`
    Type        string    `db:"type" json:"type"`
    LampCap     string    `db:"lamp_cap" json:"lamp_cap"`
    ImageURL    string    `db:"image_url" json:"image_url"`
    CreatedAt   time.Time `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type ProductRepository interface {
    GetAll() ([]Product, error)
    GetByID(id int) (Product, error)
    Create(p *Product) error
    Update(p *Product) error
    Delete(id int) error
}
