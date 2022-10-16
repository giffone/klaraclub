package db

import "github.com/jmoiron/sqlx"

type Storage interface {
}

type db struct {
	db *sqlx.DB
}

func New(base *sqlx.DB) Storage {
	return &db{
		db: base,
	}
}
