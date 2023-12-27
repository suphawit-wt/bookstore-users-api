package services

import (
	models "bookstore/users/models/users"
	"bookstore/users/utils/errors"
)

func GetUserById(userId int64) (*models.User, *errors.RestErr) {
	user := models.User{
		Id: userId,
	}
	if err := user.GetById(); err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user models.User) (*models.User, *errors.RestErr) {
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
