package main

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	port     = "5432"
	userName = "postgres"
	password = "1234590"
	host     = "localhost"
	dbname   = "claraclub_01"
)

func main() {
	connStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("postgres: connect: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("postgres: ping: %s", err)
	}

	m, err := migrate.New("file://schema", connStr)
	if err != nil {
		log.Fatalf("migrate: new: %s", err)
	}
	err = m.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migrate: up: %s", err)
	}
}
