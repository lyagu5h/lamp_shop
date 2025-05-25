-- +goose Up
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_name TEXT,
    customer_email TEXT,
    customer_phone TEXT,
    address TEXT,
    total_amount NUMERIC,
    status TEXT,
    created_at TIMESTAMP DEFAULT now()
);
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id),
    product_id INT,
    quantity INT,
    price NUMERIC
);
-- +goose Down
DROP TABLE order_items;
DROP TABLE orders;
