package models

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	DateCreated string `json:"date_created"`
}
