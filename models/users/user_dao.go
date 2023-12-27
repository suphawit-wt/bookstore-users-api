package models

import (
	"bookstore/users/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) GetById() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d Not Found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	existsUser := usersDB[user.Id]
	if existsUser != nil {
		if existsUser.Email == user.Email {
			return errors.NewConflictError(fmt.Sprintf("User Email %s Already Exists", user.Email))
		}
		return errors.NewConflictError(fmt.Sprintf("User %d Already Exists", user.Id))
	}

	usersDB[user.Id] = user
	return nil
}
