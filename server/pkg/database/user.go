package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	CreateUser(username, password string) error
	IsUserPresent(username string) (bool, error)
	GetUserPassword(username string) (string, error)
}

func (d *database) CreateUser(username, password string) error {
	query := "insert into users (username, password) values (@username, @password)"
	args := pgx.NamedArgs{
		"username": username,
		"password": password,
	}

	if _, err := d.db.Exec(context.Background(), query, args); err != nil {
		log.Println("error inserting into users: " + err.Error())
		return err
	}

	return nil
}

func (d *database) IsUserPresent(username string) (bool, error) {
	rows, err := d.db.Query(context.Background(), "select * from users where username=$1", username)
	if err != nil {
		log.Println("error fetching user: ", err)
		return false, nil
	}
	defer rows.Close()

	return rows.Next(), nil
}

func (d *database) GetUserPassword(username string) (string, error) {
	var hashedPassword string
	if err := d.db.QueryRow(context.Background(), "select password from users where username=$1", username).Scan(&hashedPassword); err != nil {
		log.Println("error fetching user: ", err)
		return "", nil
	}

	return hashedPassword, nil
}
