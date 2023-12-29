package models

import (
	"bookstore/users/database"
	"bookstore/users/logger"
	"bookstore/users/utils"
	"bookstore/users/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func (user *User) GetById() *errors.RestErr {
	query := "SELECT * FROM users WHERE id = ?;"

	err := database.DB.Get(user, query, user.Id)
	if err != nil {
		logger.Error("Database Error when trying to retrive user by id.", err)
		return utils.MySQLErrorValidate(err)
	}

	return nil
}

func (user *User) GetByStatus(status string) ([]UserPublic, *errors.RestErr) {
	query := "SELECT * FROM users WHERE status = ?;"

	users := []UserPublic{}

	err := database.DB.Select(&users, query, status)
	if err != nil {
		logger.Error("Database Error when trying to retrive users by status.", err)
		return nil, utils.MySQLErrorValidate(err)
	}

	return users, nil
}

func (user *User) Create() *errors.RestErr {
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		logger.Error("Database Error when trying to prepare insert users statement.", err)
		return errors.NewInternalServerError("Error occur on trying to insert User Data.")
	}
	defer stmt.Close()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		logger.Error("Crypto Error when trying to hash user password on insert statement.", err)
		return errors.NewInternalServerError("Error occur on trying to insert User Data.")
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, hashPass)
	if err != nil {
		logger.Error("Database Error when trying to execute insert users statement.", err)
		return utils.MySQLErrorValidate(err)
	}

	return nil
}

func (user *User) Update() *errors.RestErr {
	query := "UPDATE users SET first_name=?, last_name=?, email=?, password=? WHERE id = ?;"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		logger.Error("Database Error when trying to prepare update users statement.", err)
		return errors.NewInternalServerError("Error occur on trying to update User Data.")
	}
	defer stmt.Close()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		logger.Error("Crypto Error when trying to hash user password on update statement.", err)
		return errors.NewInternalServerError("Error occur on trying to update User Data.")
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, hashPass, user.Id)
	if err != nil {
		logger.Error("Database Error when trying to execute update users statement.", err)
		return utils.MySQLErrorValidate(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	query := "DELETE FROM users WHERE id = ?;"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		logger.Error("Database Error when trying to prepare delete user statement.", err)
		return errors.NewInternalServerError("Error occur on trying to delete User Data.")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("Database Error when trying to execute delete user statement.", err)
		return utils.MySQLErrorValidate(err)
	}

	return nil
}
