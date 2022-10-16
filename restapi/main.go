package main

import (
	"context"
	"errors"
	"fmt"
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

type storage struct {
	conn string
	pg   *sqlx.DB
}

func newDB(db *sqlx.DB) *storage {
	return &storage{
		pg: db,
	}
}

func newPg(ctx context.Context) (*storage, error) {
	var err error
	s := &storage{
		conn: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", userName, password, host, port, dbname),
	}
	s.pg, err = sqlx.ConnectContext(ctx, "postgres", s.conn)
	if err != nil {
		return nil, fmt.Errorf("postgres: connect: %s", err)
	}

	err = s.pg.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("postgres: ping: %s", err)
	}
	return s, nil
}

func (s *storage) createUser(ctx context.Context, user *User) error {
	insertSQL := `INSERT INTO "user" (f_name, l_name, age, email)
		VALUES ($1, $2, $3, $4);`

	_, err := s.pg.ExecContext(ctx, insertSQL, user.FName, user.LName, user.Age, user.Email)
	if err != nil {
		return fmt.Errorf("createUser: exec: %w", err)
	}
	return nil
}

func (s *storage) readAllUsers(ctx context.Context) ([]*User, error) {
	selectAllSQL := `SELECT id, f_name, l_name, age, email FROM "user";`
	var users []*User
	stmt, err := s.pg.PrepareContext(ctx, selectAllSQL)
	if err != nil {
		return nil, fmt.Errorf("readAllUsers: prepare statement: %w", err)
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("readAllUsers: query: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		user := &User{}
		if err = rows.Scan(&user.ID, &user.FName, &user.LName, &user.Age, &user.Email); err != nil {
			return nil, fmt.Errorf("readAllUsers: scan rows: %w", err)
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("readAllUsers: scan rows: err: %w", err)
	}
	return users, nil
}

func (s *storage) updateUserAge(ctx context.Context, age int, sort string) error {
	updateUserSQLShema := `UPDATE "user" SET age = $1 %s;`
	updateUserSQL := fmt.Sprintf(updateUserSQLShema, sort)

	stmt, err := s.pg.PrepareContext(ctx, updateUserSQL)
	if err != nil {
		return fmt.Errorf("updateUserAge: prepare statement: %w", err)
	}
	_, err = stmt.ExecContext(ctx, age)
	if err != nil {
		return fmt.Errorf("updateUserAge: exec: %w", err)
	}
	return nil
}

func (s *storage) deleteUserByID(ctx context.Context, id int) error {
	deleteUserByIDSQL := `DELETE FROM "user" WHERE id = $1;`

	stmt, err := s.pg.PrepareContext(ctx, deleteUserByIDSQL)
	if err != nil {
		return fmt.Errorf("DeleteUserByID: prepare statement: %w", err)
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("DeleteUserByID: exec: %w", err)
	}
	return nil
}

func (s *storage) printAllUsers(ctx context.Context) {
	users, err := s.readAllUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("reading users ...")
	if users == nil {
		fmt.Println("no users")
		return
	}
	for _, u := range users {
		fmt.Println(u)
	}
	fmt.Println("reading users ... done")
}

type User struct {
	ID    int    `db:"id"`
	FName string `db:"f_name"`
	LName string `db:"l_name"`
	Age   int    `db:"age"`
	Email string `db:"email"`
}

func (u User) Read() string {
	return fmt.Sprintf("{id %d; name: %s %s; age: %d; email: %s}", u.ID, u.FName, u.LName, u.Age, u.Email)
}

func main() {
	ctx := context.Background()
	db, err := newPg(ctx)
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.New("file://schema", db.conn)
	if err != nil {
		log.Fatalf("migrate: new: %s", err)
	}
	err = m.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migrate: up: %s", err)
	}

	db.printAllUsers(ctx)

	user := &User{
		FName: "Faizulla",
		LName: "Galimzhanov",
		Age:   37,
		Email: "faizulla@mail.ru",
	}
	fmt.Println("creating user")
	err = db.createUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	db.printAllUsers(ctx)

	fmt.Println("update user age")
	err = db.updateUserAge(ctx, 38, "WHERE f_name = 'Faizulla'")
	if err != nil {
		log.Fatal(err)
	}

	db.printAllUsers(ctx)

	id := 1
	fmt.Printf("delete user by id %d", id)
	db.deleteUserByID(ctx, 1)

	db.printAllUsers(ctx)
}
