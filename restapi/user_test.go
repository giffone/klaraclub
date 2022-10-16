package main

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestCRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	db2 := newDB(sqlxDB)

	u1 := User{
		FName: "test_data_01",
		LName: "test_data_01",
		Age:   1,
		Email: "test_data_01@mail.ru",
	}
	insertSQL := `INSERT INTO user (f_name, l_name, age, email)
		VALUES ($1, $2, $3, $4);`

	mock.ExpectExec(insertSQL).
		WithArgs(u1.ID, u1.FName, u1.LName, u1.Age, u1.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := db2.createUser(context.Background(), &u1); err != nil {
		log.Println(err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
