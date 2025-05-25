-- +goose Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC NOT NULL,
    power INT,
    temperature TEXT,
    type TEXT,
    lamp_cap TEXT,
    image_url TEXT,
    description TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
-- +goose Down
DROP TABLE products;
