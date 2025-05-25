package domain

import "time"

type OrderItem struct {
    ID        int     `db:"id"        json:"id"`
    OrderID   int     `db:"order_id"  json:"order_id"`
    ProductID int     `db:"product_id" json:"product_id"`
    Quantity  int     `db:"quantity"  json:"quantity"`
    Price     float64 `db:"price"     json:"price"`
}

type Order struct {
	ID            int `db:"id" json:"id"`
	CustomerName  string `db:"customer_name" json:"customer_name"`
	CustomerEmail string `db:"customer_email" json:"customer_email"`
	CustomerPhone string `db:"customer_phone" json:"customer_phone"`
	Address       string `db:"address" json:"address"`
	TotalAmount   float64    `db:"total_amount" json:"total_amount"`
	Status        string `db:"status" json:"status"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	Items         []OrderItem `json:"items"`
}

type OrderRepository interface {
	Create(o *Order) error
	List() ([]Order, error)
	GetByID(id int) (Order, error)
	UpdateStatus(id int, status string) error
}
