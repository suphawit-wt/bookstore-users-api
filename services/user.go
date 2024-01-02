package services

import (
	"bookstore/users/repository"
)

type UserService interface {
	GetUser(int64) (*repository.User, error)
	SearchUsers(string) ([]repository.UserPublic, error)
	CreateUser(repository.User) error
	UpdateUser(int64, repository.User) error
	DeleteUser(int64) error
}
