package repository

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
    ID           int    `db:"id" json:"id"`
    Username     string `db:"username" json:"username"`
    PasswordHash string `db:"password_hash" json:"-"`
    Role         string `db:"role" json:"role"`
    CreatedAt    string `db:"created_at" json:"created_at"`
    UpdatedAt    string `db:"updated_at" json:"updated_at"`
}

type UserRepo struct {
    DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
    return &UserRepo{DB: db}
}

func (r *UserRepo) Create(user *User) error {
    query := `INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)`
    _, err := r.DB.Exec(query, user.Username, user.PasswordHash, user.Role)
    return err
}

func (r *UserRepo) GetByUsername(username string) (*User, error) {
    var user User
    query := `SELECT id, username, password_hash, role, created_at, updated_at FROM users WHERE username=$1`
    if err := r.DB.Get(&user, query, username); err != nil {
        return nil, err
    }
    return &user, nil
}

