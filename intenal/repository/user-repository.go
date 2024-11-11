package repository

import (
	"github.com/R3iwan/soundScroll/intenal/models"
	"github.com/jmoiron/sqlx"
)

func RegisterUserToDB(DB *sqlx.DB, user models.RegisterUser) error {
	Query := `INSERT INTO users(username, first_name, last_name, password, age, email) VALUES (:username, :first_name, :last_name, :password, :age, :email)`
	_, err := DB.NamedExec(Query, user)
	if err != nil {
		return err
	}

	return nil
}

func LoginUserToDB(DB *sqlx.DB, user models.LoginUser) error {
	Query := `SELECT * FROM users(username, email, password) VALUES(:username, :email, :password)`
	_, err := DB.NamedExec(Query, user)
	if err != nil {
		return err
	}

	return nil
}
