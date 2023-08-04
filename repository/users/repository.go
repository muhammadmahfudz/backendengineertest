package repository_users

import (
	"Backend_Engineer_Interview_Assignment/handler/users"
	"database/sql"
	"fmt"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *PostgreSQL {
	return &PostgreSQL{db: db}
}

func (postg *PostgreSQL) Registration(payload users.Payload) error {
	stat := `INSERT INTO users (phone_number,fullname,password) VALUES ($1,$2,$3) RETURNING id`
	id := 0
	err := postg.db.QueryRow(stat, payload.PhoneNumber, payload.FullName, payload.PhoneNumber).Scan(&id)

	if err != nil {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>", err)
	}

	return nil
}

func (postg *PostgreSQL) Update(payload users.Payload) error {
	var arg []any
	stat := `UPDATE users SET`

	if payload.FullName != "" {
		stat += ` fullname = ?`
		arg = append(arg, payload.FullName)
	}

	if payload.PhoneNumber != "" {
		stat += ` phone_number = ?`
		arg = append(arg, payload.PhoneNumber)
	}

	stat += ` WHERE id = ?`
	arg = append(arg, payload.ID)

	_, err := postg.db.Exec(stat, arg)

	if err != nil {
		return err
	}

	return nil
}

func (postg *PostgreSQL) Profile(id string) *users.Payload {
	var p *users.Payload
	err := postg.db.QueryRow("SELECT phone_number, fullname FROM products WHERE id=$1",
		id).Scan(&p.PhoneNumber, &p.FullName)

	if err != nil {
		fmt.Println(err)
	}

	return p
}
