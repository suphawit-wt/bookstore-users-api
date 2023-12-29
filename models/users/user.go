package models

import "time"

type User struct {
	Id          int64     `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name" binding:"required"`
	LastName    string    `json:"last_name" db:"last_name" binding:"required"`
	Email       string    `json:"email" db:"email" binding:"required,email"`
	Status      string    `json:"status" db:"status"`
	Password    []byte    `json:"-" db:"password"`
	DateCreated time.Time `json:"date_created" db:"date_created"`
}

type UserPublic struct {
	Id          int64     `json:"id" db:"id"`
	FirstName   string    `json:"-" db:"first_name" binding:"required"`
	LastName    string    `json:"-" db:"last_name" binding:"required"`
	Email       string    `json:"-" db:"email" binding:"required,email"`
	Status      string    `json:"status" db:"status"`
	Password    []byte    `json:"-" db:"password"`
	DateCreated time.Time `json:"date_created" db:"date_created"`
}
