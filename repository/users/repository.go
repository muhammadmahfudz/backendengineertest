package repository_users

import (
	"Backend_Engineer_Interview_Assignment/handler/users"
	"database/sql"
	"errors"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *PostgreSQL {
	return &PostgreSQL{db: db}
}

func (postg *PostgreSQL) Registration(payload users.Payload) error {
	stat := `INSERT INTO users (phone_number,fullname,password) VALUES ($1,$2,$3)`
	_, err := postg.db.Exec(stat, payload.PhoneNumber, payload.FullName, payload.Password)

	if err != nil {
		return err
	}

	return nil
}

func (postg *PostgreSQL) Update(payload users.Payload) error {

	stat := `UPDATE users SET fullname = $2, phone_number = $3 WHERE id = $1`
	_, err := postg.db.Exec(stat, payload.ID, payload.FullName, payload.PhoneNumber)

	if err != nil {
		return err
	}

	return nil
}

func (postg *PostgreSQL) Profile(id string) (users.Payload, error) {
	var p users.Payload
	err := postg.db.QueryRow("SELECT phone_number, fullname FROM users WHERE id=$1",
		id).Scan(&p.PhoneNumber, &p.FullName)

	if err != nil {
		if err == sql.ErrNoRows {
			return users.Payload{}, errors.New("user not found")
		} else {
			return users.Payload{}, err
		}
	}

	return p, nil
}

func (postg *PostgreSQL) Login(phone_number string) (users.Payload, error) {
	var p users.Payload

	if err := postg.db.QueryRow("SELECT phone_number,fullname,password FROM users WHERE phone_number=$1",
		phone_number).Scan(&p.PhoneNumber, &p.FullName, &p.Password); err != nil {
		if err == sql.ErrNoRows {
			return p, errors.New("user not found")
		} else {
			return p, err
		}
	}

	return p, nil
}
