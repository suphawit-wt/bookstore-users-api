package repository

import "time"

type User struct {
	Id          int64     `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	Status      string    `db:"status"`
	Password    []byte    `db:"password"`
	DateCreated time.Time `db:"date_created"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(int64) (*User, error)
	GetByStatus(string) ([]User, error)
	Create(User) error
	Update(int64, User) error
	Delete(int64) error
}
