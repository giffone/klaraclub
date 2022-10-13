package main

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	lock sync.Mutex
	db   *sql.DB
)

func DbMutex() (*sql.DB, error) {
	lock.Lock()
	defer lock.Unlock()
	if db != nil {
		return db, nil
	}
	var err error
	file, err := os.Create("database.db")
	if err != nil {
		return nil, err
	}
	file.Close()

	db, err = sql.Open("sqlite3", file.Name())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func main() {
	db, err := DbMutex()
	if err != nil {
		panic(err)
	}
	var v int
	r := db.QueryRow("SELECT 1")
	err = r.Scan(&v)
	fmt.Println(v, err)
}
