package services

import (
	"time"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  []byte `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  []byte `json:"password" binding:"required"`
}

type UserResponse struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"date_created"`
}

type SearchUserResponse struct {
	Id          int64     `json:"id"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"date_created"`
}

type UserService interface {
	GetUser(int64) (*UserResponse, error)
	SearchUsers(string) ([]SearchUserResponse, error)
	CreateUser(CreateUserRequest) error
	UpdateUser(int64, UpdateUserRequest) error
	DeleteUser(int64) error
}
