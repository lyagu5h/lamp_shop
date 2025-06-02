-- +goose Up
ALTER TABLE order_items
    ADD CONSTRAINT fk_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE;

-- +goose Down
ALTER TABLE order_items DROP CONSTRAINT fk_order;