package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lyagu5h/lamp_backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type orderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (r *orderRepo) List() ([]domain.Order, error) {
    var orders []domain.Order
    if err := r.db.Select(&orders, "SELECT * FROM orders ORDER BY id"); err != nil {
        return nil, err
    }
    for i := range orders {
        _ = r.db.Select(&orders[i].Items, `
            SELECT id, order_id, product_id, quantity, price
              FROM order_items WHERE order_id=$1`, orders[i].ID)
        if orders[i].Items == nil {
            orders[i].Items = []domain.OrderItem{}
        }
    }
    return orders, nil
}

func (r *orderRepo) GetByID(id int) (domain.Order, error) {
	var o domain.Order

    if err := r.db.Get(&o, `
        SELECT id, customer_name, customer_email, customer_phone,
               address, total_amount, status, created_at
          FROM orders WHERE id=$1`, id); err != nil {
        return o, err         
    }

    if err := r.db.Select(&o.Items, `
        SELECT id, order_id, product_id, quantity, price
          FROM order_items WHERE order_id=$1`, id); err != nil && err != sql.ErrNoRows {
        return o, err
    }
    if o.Items == nil {   
        o.Items = []domain.OrderItem{}
    }
    return o, nil
}

func (r *orderRepo) Create(o *domain.Order) error {
	o.CreatedAt = time.Now()
	o.Status = "new"
	var total float64
	for _, it := range o.Items {
		total += it.Price * float64(it.Quantity)
	}
	o.TotalAmount = total

	if err := r.db.QueryRowx(
		`INSERT INTO orders (customer_name,customer_email,customer_phone,address,total_amount,status,created_at)
		 VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
		o.CustomerName, o.CustomerEmail, o.CustomerPhone, o.Address, o.TotalAmount, o.Status, o.CreatedAt,
	).Scan(&o.ID); err != nil {
		return err
	}
	for _, it := range o.Items {
		_, err := r.db.Exec(
			`INSERT INTO order_items (order_id,product_id,quantity,price) VALUES ($1,$2,$3,$4)`,
			o.ID, it.ProductID, it.Quantity, it.Price,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *orderRepo) UpdateStatus(id int, status string) error {
	res, err := r.db.Exec(`UPDATE orders SET status = $1 WHERE id = $2`, status, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("order not found")
	}
	return nil
}