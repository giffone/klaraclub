package users

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type UsersRequest struct {
	FNamme    string `json:"first_name"`
	LName     string `json:"last_name"`
	Interests string `json:"interests"`
}

func (u *UsersRequest) Create() (int, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return 0, err
	}

	users := &UsersDTO{
		Data: string(data),
	}

}

func (u *UsersRequest) Read(id int) error {

}

func (u *UsersRequest) Update(id int) error {

}

func (u *UsersRequest) Delete(id int) error {

}

type UsersDTO struct {
	Id   int    `db:"id"`
	Data string `db:"data"`
}
