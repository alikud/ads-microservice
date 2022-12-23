package postgres

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
}

func NewRepository(Db *pgxpool.Pool) *Repository {
	return &Repository{}
}
