package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	UserRepository
}

type database struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Database {
	return &database{
		db: db,
	}
}

