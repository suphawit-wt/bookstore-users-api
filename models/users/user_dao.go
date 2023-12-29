package models

import (
	"bookstore/users/database"
	"bookstore/users/utils"
	"bookstore/users/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func (user *User) GetById() *errors.RestErr {
	query := "SELECT * FROM users WHERE id = ?;"

	err := database.DB.Get(user, query, user.Id)
	if err != nil {
		return utils.MySQLErrorValidate(err)
	}

	return nil
}

func (user *User) GetByStatus(status string) ([]UserPublic, *errors.RestErr) {
	query := "SELECT * FROM users WHERE status = ?;"

	users := []UserPublic{}

	err := database.DB.Select(&users, query, status)
	if err != nil {
		return nil, utils.MySQLErrorValidate(err)
	}

	return users, nil
}

func (user *User) Create() *errors.RestErr {
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return errors.NewInternalServerError("Error occur on trying to insert User Data.")
	}
	defer stmt.Close()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return errors.NewInternalServerError("Error occur on trying to insert User Data.")
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, hashPass)
	if err != nil {
		return utils.MySQLErrorValidate(err)
	}

	return nil
}

func (user *User) Update() *errors.RestErr {
	query := "UPDATE users SET first_name=?, last_name=?, email=?, password=? WHERE id = ?;"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return errors.NewInternalServerError("Error occur on trying to update User Data.")
	}
	defer stmt.Close()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return errors.NewInternalServerError("Error occur on trying to update User Data.")
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, hashPass, user.Id)
	if err != nil {
		return utils.MySQLErrorValidate(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	query := "DELETE FROM users WHERE id = ?;"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return errors.NewInternalServerError("Error occur on trying to delete User Data.")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return utils.MySQLErrorValidate(err)
	}

	return nil
}
