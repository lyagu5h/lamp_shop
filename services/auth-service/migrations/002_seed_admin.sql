-- +goose Up

INSERT INTO users (username, password_hash, role)
VALUES (
    'root_admin',
    '$2a$10$wHnmv6xF9R1R4oE1VdE2IuZoVqOQVabTKBs1DseBz3t1zllDb5MeW',  -- bcrypt("password123")
    'admin'
) ON CONFLICT (username) DO NOTHING;
